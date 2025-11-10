package repository

import (
	"arukabe/core/common/constants"
	"arukabe/gen/sqlc"
	"context"
	"errors"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/google/uuid"
)

func (pr *ProviderRepository) UpdateProviderModels(ctx context.Context, providerName constants.ProviderName) error {
	dbModels, err := pr.db.AllProviderModels(ctx)
	if err != nil {
		return err
	}
	dbProviders, err := pr.db.AllProviders(ctx)
	if err != nil {
		return err
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
		return pr.updateAnthropicModels(ctx, modelsMap, providersMap)
	}
	// TODO: update other providers

	return nil
}

func (pr *ProviderRepository) updateAnthropicModels(
	ctx context.Context,
	modelsMap map[string]bool,
	providersMap map[string]uuid.UUID,
) error {
	anthropicID, ok := providersMap[string(constants.ProviderAnthropic)]
	if !ok {
		return errors.New("anthropic provider not found")
	}

	paginated, err := pr.antClient.Models.List(
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

	_, err = pr.db.SaveModels(ctx, toSave)
	return err
}
