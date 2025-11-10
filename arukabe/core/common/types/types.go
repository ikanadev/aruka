package types

import "arukabe/gen/sqlc"

type ProviderWithModels struct {
	Provider sqlc.Provider
	Models   []sqlc.Model
}
