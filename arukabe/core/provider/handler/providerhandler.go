package handler

import (
	"arukabe/core/provider/service"
	providerv1 "arukabe/gen/connect/provider/v1"
	"context"

	"connectrpc.com/connect"
)

func NewProviderHandler(service *service.ProviderService) *ProviderHandler {
	return &ProviderHandler{service}
}

type ProviderHandler struct {
	service *service.ProviderService
}

// ListProviders implements providerv1connect.ProviderServiceHandler.
func (ph *ProviderHandler) ListProviders(
	ctx context.Context,
	req *connect.Request[providerv1.ListProvidersRequest],
) (*connect.Response[providerv1.ListProvidersResponse], error) {
	resp, err := ph.service.ListProviders(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// UpdateProviderModels implements providerv1connect.ProviderServiceHandler.
func (ph *ProviderHandler) UpdateProviderModels(
	ctx context.Context,
	req *connect.Request[providerv1.UpdateProviderModelsRequest],
) (*connect.Response[providerv1.UpdateProviderModelsResponse], error) {
	resp, err := ph.service.UpdateProviderModels(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
