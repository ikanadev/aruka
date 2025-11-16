package repository

import (
	"arukabe/gen/sqlc"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type EditChatData struct {
	ChatID uuid.UUID
	Title  *string
	Prompt *string
	Pinned *bool
}

func (cr *ChatRepository) EditChat(
	ctx context.Context,
	data EditChatData,
) (sqlc.Chat, sqlc.Model, sqlc.Provider, error) {
	// Prepare parameters with nullable fields
	params := sqlc.UpdateChatParams{
		ChatID: data.ChatID,
	}

	if data.Title != nil {
		params.Title = pgtype.Text{String: *data.Title, Valid: true}
	}

	if data.Prompt != nil {
		params.Prompt = pgtype.Text{String: *data.Prompt, Valid: true}
	}

	if data.Pinned != nil {
		params.Pinned = pgtype.Bool{Bool: *data.Pinned, Valid: true}
	}

	// Update the chat
	err := cr.db.UpdateChat(ctx, params)
	if err != nil {
		return sqlc.Chat{}, sqlc.Model{}, sqlc.Provider{}, err
	}

	// Get the updated chat with model and provider info
	result, err := cr.db.GetChatProviderModel(ctx, data.ChatID)
	if err != nil {
		return sqlc.Chat{}, sqlc.Model{}, sqlc.Provider{}, err
	}

	return result.Chat, result.Model, result.Provider, nil
}
