package pg

import (
	"arukabe/domain/models"
	"arukabe/gen/sqlc"
	"context"
	"errors"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

func NewPGRepository(
	ctx context.Context,
	queries *sqlc.Queries,
	antClient *anthropic.Client,
) PGRepository {
	return PGRepository{db: queries, ctx: ctx, antClient: antClient}
}

type PGRepository struct {
	db        *sqlc.Queries
	ctx       context.Context
	antClient *anthropic.Client
}

// ListProviders implements repo.ProviderRepository.
func (p PGRepository) ListProviders(status models.ModelStatus) ([]models.Provider, error) {
	dbRows, err := p.db.ListProviders(p.ctx, sqlc.ModelStatus(status))
	if err != nil {
		return nil, err
	}
	providersMap := lo.Reduce(
		dbRows,
		func(acc map[uuid.UUID]models.Provider, el sqlc.ListProvidersRow, index int) map[uuid.UUID]models.Provider {
			if _, ok := acc[el.ID]; !ok {
				acc[el.ID] = models.Provider{
					ID:   el.ID,
					Name: el.Name,
					Models: []models.Model{},
				}
			}
			provider := acc[el.ID]
			provider.Models = append(provider.Models, models.FromDBModel(el.Model))
			acc[el.ID] = provider
			return acc
		},
		map[uuid.UUID]models.Provider{},
	)

	return lo.Values(providersMap), nil
}

// UpdateProviderModels implements repo.ProviderRepository.
func (p PGRepository) UpdateProviderModels(providerName models.ProviderName) error {
	dbModels, err := p.db.GetProviderModels(p.ctx)
	if err != nil {
		return err
	}
	dbProviders, err := p.db.GetProviders(p.ctx)
	if err != nil {
		return err
	}
	modelsMap := lo.SliceToMap(dbModels, func(el sqlc.GetProviderModelsRow) (string, bool) {
		return el.Name + el.ModelIdentifier, true
	})
	providersMap := lo.SliceToMap(dbProviders, func(el sqlc.Provider) (string, uuid.UUID) {
		return el.Name, el.ID
	})

	if providerName != models.ProviderAnthropic {
		return nil
	}
	paginated, err := p.antClient.Models.List(
		p.ctx,
		anthropic.ModelListParams{Limit: anthropic.Int(200)},
	)
	if err != nil {
		return err
	}
	anthropicModels := paginated.Data

	toSave := make([]sqlc.SaveModelsParams, 0, len(anthropicModels))
	anthropicID, ok := providersMap[models.ProviderAnthropic.String()]
	if !ok {
		return errors.New("anthropic provider not found")
	}
	for _, antModel := range anthropicModels {
		mapKey := models.ProviderAnthropic.String() + antModel.ID
		if lo.HasKey(modelsMap, mapKey) {
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

	_, err = p.db.SaveModels(p.ctx, toSave)
	return err
}
