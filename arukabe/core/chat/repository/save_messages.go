package repository

import (
	"arukabe/core/common/types"
	"arukabe/gen/sqlc"
	"context"

	"github.com/google/uuid"
)

func (cr *ChatRepository) SaveMessages(ctx context.Context, chatID uuid.UUID, messages []types.Message) error {
	toSave := make([]sqlc.SaveMessagesParams, len(messages))
	for i, msg := range messages {
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}
		content, err := msg.Content.MarshalJSON()
		if err != nil {
			return err
		}
		toSave[i] = sqlc.SaveMessagesParams{
			ID:      id,
			ChatID:  chatID,
			Role:    msg.Role,
			Content: content,
		}
	}
	_, err := cr.db.SaveMessages(ctx, toSave)
	return err
}
