package pg

import (
	"arukabe/gen/sqlc"
	"context"

	"github.com/anthropics/anthropic-sdk-go"
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
