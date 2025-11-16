package service

import (
	"arukabe/core/chat/repository"
	chatv1 "arukabe/gen/connect/chat/v1"
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

func (cs *ChatService) EditChat(
	ctx context.Context,
	req *chatv1.EditChatRequest,
) (*chatv1.EditChatResponse, error) {
	// Validate chat ID is provided
	if req.ChatId == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument,
			fmt.Errorf("chat_id is required"))
	}

	// Parse the chat ID from the request
	chatID, err := uuid.Parse(*req.ChatId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	// Prepare data for repository - directly use pointers from request
	data := repository.EditChatData{
		ChatID: chatID,
		Title:  req.Title,
		Prompt: req.Prompt,
		Pinned: req.Pinned,
	}

	// Update the chat
	dbChat, dbModel, dbProvider, err := cs.repo.EditChat(ctx, data)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Build response using helper function
	return &chatv1.EditChatResponse{
		Chat: chatToPB(dbChat, dbModel, dbProvider),
	}, nil
}
