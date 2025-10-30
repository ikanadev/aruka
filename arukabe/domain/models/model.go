package models

import (
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
