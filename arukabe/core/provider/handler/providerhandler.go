package handler

import (
	providerv1 "arukabe/gen/connect/aruka/provider/v1"
	"context"
)

func NewProviderHandler(service ProviderService) *ProviderHandler {
	return &ProviderHandler{service}
}

type ProviderHandler struct {
	service ProviderService
}

// ListProviders implements providerv1connect.ProviderServiceHandler.
func (ph *ProviderHandler) ListProviders(
	ctx context.Context,
	req *providerv1.ListProvidersRequest,
) (*providerv1.ListProvidersResponse, error) {
	resp, err := ph.service.ListProviders(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateProviderModels implements providerv1connect.ProviderServiceHandler.
func (ph *ProviderHandler) UpdateProviderModels(
	ctx context.Context,
	req *providerv1.UpdateProviderModelsRequest,
) (*providerv1.UpdateProviderModelsResponse, error) {
	resp, err := ph.service.UpdateProviderModels(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
