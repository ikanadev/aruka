package service

import (
	"arukabe/core/common/mappers"
	"arukabe/core/common/types"
	modelsv1 "arukabe/gen/connect/models/v1"
	providerv1 "arukabe/gen/connect/provider/v1"
	"context"
)

func (ps *ProviderService) ListProviders(
	ctx context.Context,
	req *providerv1.ListProvidersRequest,
) (*providerv1.ListProvidersResponse, error) {
	dbProviders, err := ps.repo.ListProviders(ctx, mappers.ModelStatusToDB(req.Status))
	if err != nil {
		return nil, err
	}
	resp := providerv1.ListProvidersResponse{
		Providers: mapProvidersToProto(dbProviders),
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
