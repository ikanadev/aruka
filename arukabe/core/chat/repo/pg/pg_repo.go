package pg

import (
	"arukabe/core/chat/repo"
	"arukabe/domain/models"
	"arukabe/gen/sqlc"
	"context"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/google/uuid"
)

func NewPGRepository(
	ctx context.Context,
	queries *sqlc.Queries,
	antClient *anthropic.Client,
) PGRepository {
	return PGRepository{db: queries, ctx: ctx, antClient: antClient}
}

type PGRepository struct {
	db        *sqlc.Queries
	ctx       context.Context
	antClient *anthropic.Client
}

// NewChat implements repo.ChatRepository.
func (p PGRepository) NewChat(data repo.NewChatData) (models.Chat, error) {
  id, err := uuid.NewV7()
  if err != nil {
    return models.Chat{}, err
  }

  chat, err := p.db.CreateChat(p.ctx, sqlc.CreateChatParams{
    ID: id,
    Title: "",
    Prompt: data.Prompt,
    ModelID: data.ModelID,
  })
  if err != nil {
    return models.Chat{}, err
  }
  return models.Chat{
  	ID:       id,
  	Title:    chat.Title,
  	Prompt:   chat.Prompt,
  	Pinned:   chat.Pinned,
  	Model:    models.FromDBModel(chat.Model),
  	TimeData: models.TimeData{
  		CreatedAt:  chat.CreatedAt,
  		UpdatedAt:  chat.UpdatedAt,
  		ArchivedAt: chat.ArchivedAt,
  		DeletedAt:  chat.DeletedAt,
  	},
  }, nil
}

