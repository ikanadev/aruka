package service

import (
	"arukabe/core/common/mappers"
	chatv1 "arukabe/gen/connect/aruka/chat/v1"
	"arukabe/gen/sqlc"
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

func (s *ChatService) ChatMessages(ctx context.Context, req *chatv1.ChatMessagesRequest) (*chatv1.ChatMessagesResponse, error) {
	uuid, err := uuid.Parse(req.ChatId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	dbResult, err := s.db.GetChatMessages(ctx, uuid)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	dbMessages := make([]sqlc.Message, len(dbResult))

	for i, dbRow := range dbResult {
		dbMessages[i] = dbRow.Message
	}

	messages, err := mappers.DBMessagesToPB(dbMessages)
	if err != nil {
		return nil, err
	}

	return &chatv1.ChatMessagesResponse{
		Messages: messages,
	}, nil
}
