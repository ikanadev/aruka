package repository

import (
	"arukabe/gen/sqlc"
	"context"

	"github.com/google/uuid"
)

func (cr *ChatRepository) GetChatProviderModel(ctx context.Context, chatID string) (sqlc.Chat, sqlc.Provider, sqlc.Model, error) {
	uuid, err := uuid.Parse(chatID)
	if err != nil {
		return sqlc.Chat{}, sqlc.Provider{}, sqlc.Model{}, err
	}
	dbData, err := cr.db.GetChatProviderModel(ctx, uuid)
	if err != nil {
		return sqlc.Chat{}, sqlc.Provider{}, sqlc.Model{}, err
	}
	return dbData.Chat, dbData.Provider, dbData.Model, nil
}
