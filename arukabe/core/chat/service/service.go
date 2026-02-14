package service

import (
	"arukabe/gen/sqlc"

	"github.com/anthropics/anthropic-sdk-go"
)

type ChatService struct {
	db *sqlc.Queries
	antClient *anthropic.Client
}

func NewChatService(db *sqlc.Queries, antClient *anthropic.Client) *ChatService {
	return &ChatService{db, antClient}
}
