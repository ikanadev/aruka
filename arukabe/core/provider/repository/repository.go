package repository

import (
	"arukabe/gen/sqlc"

	"github.com/anthropics/anthropic-sdk-go"
)

type ProviderRepository struct {
	db *sqlc.Queries
	antClient *anthropic.Client
}

func NewProviderRepository(db *sqlc.Queries, antClient *anthropic.Client) *ProviderRepository {
	return &ProviderRepository{db: db, antClient: antClient}
}
