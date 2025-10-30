package handler

import (
	"arukabe/core/provider/repo"
	providerv1 "arukabe/gen/connect/provider/v1"
	"context"

	"connectrpc.com/connect"
)

func NewProviderHandler(repo repo.ProviderRepository) *ProviderHandler {
	return &ProviderHandler{repo}
}

type ProviderHandler struct {
	repo repo.ProviderRepository
}

// ListProviders implements providerv1connect.ProviderServiceHandler.
func (p *ProviderHandler) ListProviders(
	ctx context.Context,
	req *connect.Request[providerv1.ListProvidersRequest],
) (*connect.Response[providerv1.ListProvidersResponse], error) {
	status := toModelStatus(req.Msg.Status)
	providers, err := p.repo.ListProviders(status)
	if err != nil {
		return nil, err
	}
	resp := &providerv1.ListProvidersResponse{
		Providers: fromProviders(providers),
	}
	return connect.NewResponse(resp), nil
}

// UpdateProviderModels implements providerv1connect.ProviderServiceHandler.
func (p *ProviderHandler) UpdateProviderModels(
	ctx context.Context,
	req *connect.Request[providerv1.UpdateProviderModelsRequest],
) (*connect.Response[providerv1.UpdateProviderModelsResponse], error) {
	providerName := toProviderName(req.Msg.ProviderName)
	err := p.repo.UpdateProviderModels(providerName)
	return connect.NewResponse(&providerv1.UpdateProviderModelsResponse{}), err
}
