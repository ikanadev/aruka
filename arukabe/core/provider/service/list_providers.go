package service

import (
	"arukabe/core/common/mappers"
	"arukabe/core/common/types"
	modelsv1 "arukabe/gen/connect/aruka/models/v1"
	providerv1 "arukabe/gen/connect/aruka/provider/v1"
	"arukabe/gen/sqlc"
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

func (ps *ProviderService) ListProviders(
	ctx context.Context,
	req *providerv1.ListProvidersRequest,
) (*providerv1.ListProvidersResponse, error) {
	status := mappers.ModelStatusToDB(req.Status)
	dbRows, err := ps.db.ListProviders(ctx, status)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	providerMap := make(map[uuid.UUID]types.ProviderWithModels)
	for _, row := range dbRows {
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

	resp := providerv1.ListProvidersResponse{
		Providers: mapProvidersToProto(providers),
	}
	return &resp, nil
}

func mapProvidersToProto(dbProviders []types.ProviderWithModels) []*modelsv1.Provider {
	providers := make([]*modelsv1.Provider, len(dbProviders))
	for i := range dbProviders {
		provider := dbProviders[i]
		providers[i] = &modelsv1.Provider{
			Id:   provider.Provider.ID.String(),
			Name: provider.Provider.Name,
		}
		providers[i].Models = mappers.ModelsToPB(provider.Models)
	}
	return providers
}
