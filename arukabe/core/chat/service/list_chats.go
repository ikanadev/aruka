package service

import (
	"arukabe/core/common/mappers"
	chatv1 "arukabe/gen/connect/chat/v1"
	modelsv1 "arukabe/gen/connect/models/v1"
	"arukabe/gen/sqlc"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ChatService) ListChats(ctx context.Context, req *chatv1.ListChatsRequest) (*chatv1.ListChatsResponse, error) {
	dbChats, err := s.repo.GetAllChats(ctx)
	if err != nil {
		return nil, err
	}

	// Group rows by chat ID to handle potential duplicates
	chatMap := make(map[string]*modelsv1.Chat)
	for _, row := range dbChats {
		chatID := row.Chat.ID.String()
		if _, exists := chatMap[chatID]; !exists {
			chatMap[chatID] = chatToPB(row.Chat, row.Model, row.Provider)
		}
	}

	// Convert map to slice
	chats := make([]*modelsv1.Chat, 0, len(chatMap))
	for _, chat := range chatMap {
		chats = append(chats, chat)
	}

	return &chatv1.ListChatsResponse{
		Chats: chats,
		// TODO: Implement pagination if needed
		Pagination: nil,
	}, nil
}

func chatToPB(chat sqlc.Chat, model sqlc.Model, provider sqlc.Provider) *modelsv1.Chat {
	return &modelsv1.Chat{
		Id:     chat.ID.String(),
		Title:  chat.Title,
		Prompt: chat.Prompt,
		Pinned: chat.Pinned,
		Model:  mappers.ModelToPB(model),
		Provider: &modelsv1.ProviderBase{
			Id:   provider.ID.String(),
			Name: provider.Name,
		},
		TimeData: &modelsv1.TimeData{
			CreatedAt:  timestamppb.New(chat.CreatedAt),
			UpdatedAt:  timestamppb.New(chat.UpdatedAt),
			ArchivedAt: mappers.TimestampFromTime(chat.ArchivedAt),
			DeletedAt:  mappers.TimestampFromTime(chat.DeletedAt),
		},
	}
}
