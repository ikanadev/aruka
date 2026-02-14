package service

import (
	"arukabe/gen/sqlc"

	"github.com/anthropics/anthropic-sdk-go"
)

type ProviderService struct {
	db *sqlc.Queries
	antClient *anthropic.Client
}

func NewProviderService(db *sqlc.Queries, antClient *anthropic.Client) *ProviderService {
	return &ProviderService{db, antClient}
}
