package service

import (
	"arukabe/core/chat/repository"
	"arukabe/core/common/constants"
	"arukabe/core/common/mappers"
	"arukabe/core/common/types"
	chatv1 "arukabe/gen/connect/chat/v1"
	"arukabe/gen/sqlc"
	"context"
	"fmt"
	"sync"
)

func (cs *ChatService) ChatMessage(
	ctx context.Context,
	req *chatv1.ChatMessageRequest,
) (<-chan types.ChatStreamResult, error) {
	dbChat, dbProvider, dbModel, dbMessages, err := getChatData(ctx, req, cs.repo)
	if err != nil {
		return nil, err
	}

	var messageContent types.MessageSections
	if err := messageContent.FromPBSections(req.Content); err != nil {
		return nil, err
	}
	userMessage := types.Message{
		Role:    sqlc.MessageRoleUSER,
		Content: messageContent,
	}
	messages, err := mappers.FromDBMessages(dbMessages)
	if err != nil {
		return nil, err
	}
	messages = append(messages, userMessage)

	if dbProvider.Name == string(constants.ProviderAnthropic) {
		return cs.repo.HandleAnthropicChatMessage(ctx, dbChat, dbModel, messages)
	}

	if dbProvider.Name == string(constants.ProviderOpenAI) {
		return nil, fmt.Errorf("open AI provider not available: %s", dbProvider.Name)
	}
	return nil, fmt.Errorf("unknown provider: %s", dbProvider.Name)
}

func getChatData(
	ctx context.Context,
	req *chatv1.ChatMessageRequest,
	repo *repository.ChatRepository,
) (sqlc.Chat, sqlc.Provider, sqlc.Model, []sqlc.Message, error) {
	var dbChat sqlc.Chat
	var dbProvider sqlc.Provider
	var dbModel sqlc.Model
	var dbMessages []sqlc.Message
	var getChatProviderModelErr, getChatMessagesErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		dbChat, dbProvider, dbModel, getChatProviderModelErr = repo.
			GetChatProviderModel(ctx, req.ChatId)
	}()
	go func() {
		defer wg.Done()
		dbMessages, getChatMessagesErr = repo.GetChatMessages(ctx, req.ChatId)
	}()
	wg.Wait()
	if getChatProviderModelErr != nil {
		return dbChat, dbProvider, dbModel, dbMessages, fmt.Errorf("getChatProviderModelErr: %w", getChatProviderModelErr)
	}
	if getChatMessagesErr != nil {
		return dbChat, dbProvider, dbModel, dbMessages, fmt.Errorf("getChatMessagesErr: %w", getChatMessagesErr)
	}
	return dbChat, dbProvider, dbModel, dbMessages, nil
}
