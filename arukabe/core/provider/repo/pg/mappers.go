package pg

import (
	"arukabe/domain/models"
	"arukabe/gen/sqlc"
)

func fromDBModel(row sqlc.ListProvidersRow) models.Model {
	return models.Model{
		ID:              row.ModelID,
		Name:            row.ModelName,
		ModelIdentifier: row.ModelIdentifier,
		Status:          models.ModelStatus(row.ModelStatus),
		CreatedAt:       row.ModelCreatedAt,
	}
}
