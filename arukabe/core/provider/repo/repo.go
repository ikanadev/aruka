package repo

import "arukabe/domain/models"

type ProviderRepository interface {
	UpdateProviderModels(providerName models.ProviderName) error
	ListProviders(status models.ModelStatus) ([]models.Provider, error)
}
