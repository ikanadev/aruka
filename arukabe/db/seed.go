package db

import (
	"arukabe/core/common/constants"
	"arukabe/gen/sqlc"
	"context"

	"github.com/google/uuid"
)

func SeedProviders(ctx context.Context, queries *sqlc.Queries) error {
	dbProviders, err := queries.AllProviders(ctx)
	if err != nil {
		return err
	}
	namesToSave := make([]sqlc.SaveProvidersParams, 0, len(constants.ProviderNames))

	// filter names with existing dbNames
	for _, name := range constants.ProviderNames {
		exists := false
		for _, dbProvider := range dbProviders {
			if string(name) == dbProvider.Name {
				exists = true
				break
			}
		}
		if exists {
			continue
		}
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}
		namesToSave = append(namesToSave, sqlc.SaveProvidersParams{
			ID:   id,
			Name: string(name),
		})
	}

	_, err = queries.SaveProviders(ctx, namesToSave)
	return err
}
