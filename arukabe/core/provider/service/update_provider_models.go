package service

import (
	"arukabe/core/common/constants"
	providerv1 "arukabe/gen/connect/aruka/provider/v1"
	"arukabe/gen/sqlc"
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/google/uuid"
)

func (ps *ProviderService) UpdateProviderModels(
	ctx context.Context,
	req *providerv1.UpdateProviderModelsRequest,
) (*providerv1.UpdateProviderModelsResponse, error) {
	providerName, err := constants.ToProviderName(req.ProviderName)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	dbModels, err := ps.db.AllProviderModels(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}
	dbProviders, err := ps.db.AllProviders(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	modelsMap := make(map[string]bool, len(dbModels))
	for _, dbModel := range dbModels {
		modelsMap[dbModel.Provider.Name+dbModel.Model.ModelIdentifier] = true
	}
	providersMap := make(map[string]uuid.UUID, len(dbProviders))
	for _, prov := range dbProviders {
		providersMap[prov.Name] = prov.ID
	}

	if providerName == constants.ProviderAnthropic {
		err = ps.updateAnthropicModels(ctx, modelsMap, providersMap)
		return &providerv1.UpdateProviderModelsResponse{}, err
	}
	// TODO: update other providers

	return &providerv1.UpdateProviderModelsResponse{}, err
}

func (ps *ProviderService) updateAnthropicModels(
	ctx context.Context,
	modelsMap map[string]bool,
	providersMap map[string]uuid.UUID,
) error {
	anthropicID, ok := providersMap[string(constants.ProviderAnthropic)]
	if !ok {
		return errors.New("anthropic provider not found")
	}

	paginated, err := ps.antClient.Models.List(
		ctx,
		anthropic.ModelListParams{Limit: anthropic.Int(200)},
	)
	if err != nil {
		return err
	}
	anthropicModels := paginated.Data

	toSave := make([]sqlc.SaveModelsParams, 0, len(anthropicModels))
	for _, antModel := range anthropicModels {
		mapKey := string(constants.ProviderAnthropic) + antModel.ID
		if _, ok := modelsMap[mapKey]; ok {
			continue
		}
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}
		toSave = append(toSave, sqlc.SaveModelsParams{
			ID:              id,
			Name:            antModel.DisplayName,
			ProviderID:      anthropicID,
			ModelIdentifier: antModel.ID,
			CreatedAt:       antModel.CreatedAt,
		})
	}

	_, err = ps.db.SaveModels(ctx, toSave)
	return err
}
