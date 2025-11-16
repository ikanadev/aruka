package repository

import (
	"arukabe/gen/sqlc"
	"context"

	"github.com/google/uuid"
)

func (cr *ChatRepository) GetChatMessages(ctx context.Context, chatID string) ([]sqlc.Message, error) {
	uuid, err := uuid.Parse(chatID)
	if err != nil {
		return nil, err
	}
	dbData, err := cr.db.GetChatMessages(ctx, uuid)
	if err != nil {
		return nil, err
	}
	messages := make([]sqlc.Message, len(dbData))
	for i, msg := range dbData {
		messages[i] = msg.Message
	}
	return messages, nil
}
