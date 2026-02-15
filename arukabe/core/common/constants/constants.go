package constants

import "errors"

type ProviderName string

const (
	ProviderOpenAI    ProviderName = "OpenAI"
	ProviderAnthropic ProviderName = "Anthropic"
	ProviderGoogle    ProviderName = "Google"
	ProviderMeta      ProviderName = "Meta"
)

const TitleGenerationModel = "claude-3-5-haiku-20241022"

func ToProviderName(name string) (ProviderName, error) {
	switch name {
	case string(ProviderOpenAI):
		return ProviderOpenAI, nil
	case string(ProviderAnthropic):
		return ProviderAnthropic, nil
	case string(ProviderGoogle):
		return ProviderGoogle, nil
	case string(ProviderMeta):
		return ProviderMeta, nil
	default:
		return "", errors.New("invalid provider name")
	}
}

var ProviderNames = [4]ProviderName{
	ProviderOpenAI,
	ProviderAnthropic,
	ProviderGoogle,
	ProviderMeta,
}
