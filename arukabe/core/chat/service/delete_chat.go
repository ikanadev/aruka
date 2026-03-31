package service

import (
	chatv1 "arukabe/gen/connect/aruka/chat/v1"
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

func (cs *ChatService) DeleteChat(
	ctx context.Context,
	req *chatv1.DeleteChatRequest,
) (*chatv1.DeleteChatResponse, error) {
	chatID, err := uuid.Parse(req.ChatId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	if err := cs.db.DeleteChat(ctx, chatID); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &chatv1.DeleteChatResponse{}, nil
}
