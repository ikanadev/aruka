package db

import (
	"arukabe/domain/models"
	"arukabe/gen/sqlc"
	"context"

	"github.com/google/uuid"
)

func SeedProviders(ctx context.Context, queries *sqlc.Queries) error {
	dbProviders, err := queries.GetProviders(ctx)
	if err != nil {
		return err
	}
	namesToSave := make([]models.ProviderName, 0, len(models.ProviderNames))

	// filter names with existing dbNames
	for _, name := range models.ProviderNames {
		exists := false
		for _, dbProvider := range dbProviders {
			if name.String() == dbProvider.Name {
				exists = true
				break
			}
		}
		if !exists {
			namesToSave = append(namesToSave, name)
		}
	}

	toSave := make([]sqlc.SaveProvidersParams, len(namesToSave))
	for i, name := range namesToSave {
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}
		toSave[i] = sqlc.SaveProvidersParams{
			ID:   id,
			Name: name.String(),
		}
	}
	_, err = queries.SaveProviders(ctx, toSave)
	return err
}
