package repository

import (
	"arukabe/core/common/types"
	"arukabe/gen/sqlc"
	"context"

	"github.com/google/uuid"
)

func (pr *ProviderRepository) ListProviders(
	ctx context.Context,
	status sqlc.ModelStatus,
) ([]types.ProviderWithModels, error) {
	rows, err := pr.db.ListProviders(ctx, status)
	if err != nil {
		return nil, err
	}
	providerMap := make(map[uuid.UUID]types.ProviderWithModels)

	for _, row := range rows {
		if _, ok := providerMap[row.Provider.ID]; !ok {
			providerMap[row.Provider.ID] = types.ProviderWithModels{
				Provider: sqlc.Provider{
					ID:   row.Provider.ID,
					Name: row.Provider.Name,
				},
				Models: []sqlc.Model{},
			}
		}

		provider := providerMap[row.Provider.ID]
		provider.Models = append(provider.Models, row.Model)
		providerMap[row.Provider.ID] = provider
	}

	providers := make([]types.ProviderWithModels, 0, len(providerMap))
	for _, v := range providerMap {
		providers = append(providers, v)
	}

	return providers, err
}
