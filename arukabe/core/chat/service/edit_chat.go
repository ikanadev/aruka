package service

import (
	chatv1 "arukabe/gen/connect/aruka/chat/v1"
	"arukabe/gen/sqlc"
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
	data := sqlc.UpdateChatParams{
		ChatID: chatID,
	}
	if req.Title != nil {
		data.Title = pgtype.Text{String: *req.Title, Valid: true}
	}
	if req.Prompt != nil {
		data.Prompt = pgtype.Text{String: *req.Prompt, Valid: true}
	}

	if req.Pinned != nil {
		data.Pinned = pgtype.Bool{Bool: *req.Pinned, Valid: true}
	}

	// Update the chat
	err = cs.db.UpdateChat(ctx, data)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Get the updated chat with model and provider info
	dbRow, err := cs.db.GetChatProviderModel(ctx, data.ChatID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Build response using helper function
	return &chatv1.EditChatResponse{
		Chat: chatToPB(dbRow.Chat, dbRow.Model, dbRow.Provider),
	}, nil
}
