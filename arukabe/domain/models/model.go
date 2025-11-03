package models

import (
	"arukabe/gen/sqlc"
	"time"

	"github.com/google/uuid"
)

type ModelStatus string

const (
	ModelStatusActive     ModelStatus = "ACTIVE"
	ModelStatusInactive   ModelStatus = "INACTIVE"
	ModelStatusDeprecated ModelStatus = "DEPRECATED"
)

type Model struct {
	ID              uuid.UUID
	ModelIdentifier string
	Name            string
	Status          ModelStatus
	CreatedAt       time.Time
}

func FromDBModel(row sqlc.Model) Model {
	return Model{
		ID:              row.ID,
		Name:            row.Name,
		ModelIdentifier: row.ModelIdentifier,
		Status:          ModelStatus(row.Status),
		CreatedAt:       row.CreatedAt,
	}
}
