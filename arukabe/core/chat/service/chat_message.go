package service

import (
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
) (<-chan string, <-chan error, error) {
	var dbChat sqlc.Chat
	var dbProvider sqlc.Provider
	var dbModel sqlc.Model
	var dbMessages []sqlc.Message
	var getChatProviderModelErr, getChatMessagesErr error

	/*
		dbChat, dbProvider, dbModel, getChatProviderModelErr = cs.repo.GetChatProviderModel(ctx, req.ChatId)
		dbMessages, getChatMessagesErr = cs.repo.GetChatMessages(ctx, req.ChatId)
	*/
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		dbChat, dbProvider, dbModel, getChatProviderModelErr = cs.
			repo.
			GetChatProviderModel(ctx, req.ChatId)
	}()
	go func() {
		defer wg.Done()
		dbMessages, getChatMessagesErr = cs.repo.GetChatMessages(ctx, req.ChatId)
	}()
	wg.Wait()

	if getChatProviderModelErr != nil {
		fmt.Println("getChatProviderModelErr: ", getChatProviderModelErr)
		return nil, nil, getChatProviderModelErr
	}
	if getChatMessagesErr != nil {
		fmt.Println("getChatMessagesErr: ", getChatMessagesErr)
		return nil, nil, getChatMessagesErr
	}

	messageContent := mappers.PBMessageContentToMessageSections(req.Content)
	userMessage := types.Message{
		Role:    sqlc.MessageRoleUSER,
		Content: messageContent,
	}
	messages, err := mappers.FromDBMessages(dbMessages)
	if err != nil {
		fmt.Println("mappers errors: ", getChatMessagesErr)
		return nil, nil, err
	}
	messages = append(messages, userMessage)

	if dbProvider.Name == string(constants.ProviderAnthropic) {
		return cs.repo.HandleAnthropicChatMessage(ctx, dbChat, dbModel, messages)
	}

	if dbProvider.Name == string(constants.ProviderOpenAI) {
		return nil, nil, fmt.Errorf("open AI provider not available: %s", dbProvider.Name)
	}
	return nil, nil, fmt.Errorf("unknown provider: %s", dbProvider.Name)

	// processedTextChan := make(chan string)
	// processedErrChan := make(chan error, 1)
	//
	// go func() {
	//     defer close(processedTextChan)
	//     defer close(processedErrChan)
	//
	//     for {
	//         select {
	//         case text, ok := <-textChan:
	//             if !ok {
	//                 return
	//             }
	//             // Apply business logic transformation
	//             processedText := cs.processText(text)
	//             processedTextChan <- processedText
	//         case err := <-errChan:
	//             processedErrChan <- err
	//             return
	//         case <-ctx.Done():
	//             processedErrChan <- ctx.Err()
	//             return
	//         }
	//     }
	// }()
	//
	// return processedTextChan, processedErrChan
}
