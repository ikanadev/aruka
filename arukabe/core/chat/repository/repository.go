package repository

import (
	"arukabe/gen/sqlc"

	"github.com/anthropics/anthropic-sdk-go"
)

type ChatRepository struct {
	db *sqlc.Queries
	antClient *anthropic.Client
}

func NewChatRepository(db *sqlc.Queries, antClient *anthropic.Client) *ChatRepository {
	return &ChatRepository{db: db, antClient: antClient}
}
