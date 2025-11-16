package service

import (
	"arukabe/core/common/mappers"
	chatv1 "arukabe/gen/connect/chat/v1"
	"context"
)

func (s *ChatService) ChatMessages(ctx context.Context, req *chatv1.ChatMessagesRequest) (*chatv1.ChatMessagesResponse, error) {
	dbMessages, err := s.repo.GetChatMessages(ctx, req.ChatId)
	if err != nil {
		return nil, err
	}

	messages, err := mappers.DBMessagesToPB(dbMessages)
	if err != nil {
		return nil, err
	}

	return &chatv1.ChatMessagesResponse{
		Messages: messages,
	}, nil
}
