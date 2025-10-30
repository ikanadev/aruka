package models

import "github.com/google/uuid"

type ProviderName string

const (
	ProviderOpenAI    ProviderName = "OpenAI"
	ProviderAnthropic ProviderName = "Anthropic"
	ProviderGoogle    ProviderName = "Google"
	ProviderMeta      ProviderName = "Meta"
)

func (provider ProviderName) String() string {
	return string(provider)
}

var ProviderNames = [4]ProviderName{
	ProviderOpenAI,
	ProviderAnthropic,
	ProviderGoogle,
	ProviderMeta,
}

type Provider struct {
	ID     uuid.UUID
	Name   string
	Models []Model
}
