package service

import (
	"arukabe/core/chat/service/anthropicmsg"
	"arukabe/core/common/constants"
	"arukabe/core/common/mappers"
	"arukabe/core/common/types"
	chatv1 "arukabe/gen/connect/aruka/chat/v1"
	"arukabe/gen/sqlc"
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func (cs *ChatService) ChatMessage(
	ctx context.Context,
	req *chatv1.ChatMessageRequest,
) (<-chan types.ChatStreamResult, error) {
	dbChat, dbProvider, dbModel, dbMessages, err := getChatData(ctx, req.ChatId, cs.db)
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
		return anthropicmsg.HandleAnthropicChatMessage(ctx, dbChat, dbModel, messages, cs.antClient, cs.db)
	}

	if dbProvider.Name == string(constants.ProviderOpenAI) {
		return nil, fmt.Errorf("open AI provider not available: %s", dbProvider.Name)
	}
	return nil, fmt.Errorf("unknown provider: %s", dbProvider.Name)
}

func getChatData(
	ctx context.Context,
	chatId string,
	db *sqlc.Queries,
) (sqlc.Chat, sqlc.Provider, sqlc.Model, []sqlc.Message, error) {
	var dbChat sqlc.Chat
	var dbProvider sqlc.Provider
	var dbModel sqlc.Model
	var getChatProviderModelErr, getChatMessagesErr error
	dbMessages := make([]sqlc.Message, 0)
	uuid, err := uuid.Parse(chatId)

	if err != nil {
		return dbChat, dbProvider, dbModel, dbMessages, err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		dbData, err := db.GetChatProviderModel(ctx, uuid)
		getChatProviderModelErr = err
		dbChat = dbData.Chat
		dbProvider = dbData.Provider
		dbModel = dbData.Model
	}()
	go func() {
		defer wg.Done()
		dbData, err := db.GetChatMessages(ctx, uuid)
		getChatMessagesErr = err
		for _, msg := range dbData {
			dbMessages = append(dbMessages, msg.Message)
		}
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
