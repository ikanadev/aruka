package handler

import (
	providerv1 "arukabe/gen/connect/aruka/provider/v1"
	"context"
)

type ProviderService interface {
	ListProviders(ctx context.Context, req *providerv1.ListProvidersRequest) (*providerv1.ListProvidersResponse, error)
	UpdateProviderModels(ctx context.Context, req *providerv1.UpdateProviderModelsRequest) (*providerv1.UpdateProviderModelsResponse, error)
}
