package repository

import (
	"arukabe/gen/sqlc"
	"context"
	"strings"

	"github.com/google/uuid"
)

type NewChatData struct {
	Prompt  string
	ModelID uuid.UUID
}
func (cr *ChatRepository) NewChat(ctx context.Context, data NewChatData) (sqlc.Chat, sqlc.Model, sqlc.Provider, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return sqlc.Chat{}, sqlc.Model{}, sqlc.Provider{}, err
	}
	dbData, err := cr.db.CreateChat(ctx, sqlc.CreateChatParams{
		ID:       id,
		Title:    "",
		Prompt:   strings.TrimSpace(data.Prompt),
		ModelID:  data.ModelID,
	})
	if err != nil {
		return sqlc.Chat{}, sqlc.Model{}, sqlc.Provider{}, err
	}
	chat := sqlc.Chat{
		ID:       dbData.ID,
		Title:    dbData.Title,
		Prompt:   dbData.Prompt,
		Pinned:   dbData.Pinned,
		ModelID:  dbData.ModelID,
		CreatedAt: dbData.CreatedAt,
		UpdatedAt: dbData.UpdatedAt,
		ArchivedAt: dbData.ArchivedAt,
		DeletedAt:  dbData.DeletedAt,
	}
	return chat, dbData.Model, dbData.Provider, nil
}
