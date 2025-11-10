package service

import (
	"arukabe/core/common/constants"
	providerv1 "arukabe/gen/connect/provider/v1"
	"context"

	"connectrpc.com/connect"
)

func (ps *ProviderService) UpdateProviderModels(
	ctx context.Context,
	req *providerv1.UpdateProviderModelsRequest,
) (*providerv1.UpdateProviderModelsResponse, error) {
	providerName, err := constants.ToProviderName(req.ProviderName)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	err = ps.repo.UpdateProviderModels(ctx, providerName)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &providerv1.UpdateProviderModelsResponse{}, err
}
