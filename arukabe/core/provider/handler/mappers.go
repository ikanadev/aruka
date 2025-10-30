package handler

import (
	"arukabe/domain/models"
	modelsv1 "arukabe/gen/connect/models/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toModelStatus(status modelsv1.ModelStatus) models.ModelStatus {
	// swtich case in go
	switch status {
	case modelsv1.ModelStatus_MODEL_STATUS_UNSPECIFIED:
	case modelsv1.ModelStatus_MODEL_STATUS_ACTIVE:
		return models.ModelStatusActive
	case modelsv1.ModelStatus_MODEL_STATUS_INACTIVE:
		return models.ModelStatusInactive
	case modelsv1.ModelStatus_MODEL_STATUS_DEPRECATED:
		return models.ModelStatusDeprecated
	default:
		return models.ModelStatusActive
	}
	return models.ModelStatusActive
}

func fromModelStatus(status models.ModelStatus) modelsv1.ModelStatus {
	switch status {
	case models.ModelStatusActive:
		return modelsv1.ModelStatus_MODEL_STATUS_ACTIVE
	case models.ModelStatusInactive:
		return modelsv1.ModelStatus_MODEL_STATUS_INACTIVE
	case models.ModelStatusDeprecated:
		return modelsv1.ModelStatus_MODEL_STATUS_DEPRECATED
	default:
		return modelsv1.ModelStatus_MODEL_STATUS_ACTIVE
	}
}

func fromProviders(providers []models.Provider) []*modelsv1.Provider {
	result := make([]*modelsv1.Provider, len(providers))
	for i := range providers {
		provider := providers[i]

		models := make([]*modelsv1.Model, len(provider.Models))
		for j := range provider.Models {
			model := provider.Models[j]
			models[j] = &modelsv1.Model{
				Id:   model.ID.String(),
				Name: model.Name,
				ModelIdentifier: model.ModelIdentifier,
				CreatedAt: timestamppb.New(model.CreatedAt),
				Status:  fromModelStatus(model.Status),
			}
		}

		result[i] = &modelsv1.Provider{
			Id:   provider.ID.String(),
			Name: provider.Name,
			Models: models,
		}
	}
	return result
}

func toProviderName(name string) models.ProviderName {
	switch name {
	case models.ProviderOpenAI.String():
		return models.ProviderOpenAI
	case models.ProviderAnthropic.String():
		return models.ProviderAnthropic
	case models.ProviderGoogle.String():
		return models.ProviderGoogle
	case models.ProviderMeta.String():
		return models.ProviderMeta
	default:
		return models.ProviderOpenAI
	}
}
