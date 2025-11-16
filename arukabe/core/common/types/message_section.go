package types

import (
	"encoding/json"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

type MessageSectionType string

const (
	MessageSectionTypeText MessageSectionType = "text"
)

type MessageTextSection struct {
	Text string `json:"text"`
}

type MessageSection struct {
	Type    MessageSectionType `json:"type"`
	Payload any                `json:"payload"`
}
type MessageSections []MessageSection

func (m *MessageSections) ToAnthropicMessageSections() ([]anthropic.ContentBlockParamUnion, error) {
	var sections []anthropic.ContentBlockParamUnion
	for _, section := range *m {
		switch section.Type {
		case MessageSectionTypeText:
			textContent := section.Payload.(MessageTextSection)
			sections = append(sections, anthropic.ContentBlockParamUnion{
				OfText: &anthropic.TextBlockParam{Text: textContent.Text},
			})
		// TODO: handle more content types
		default:
			return nil, fmt.Errorf("unknown message content type: %s", section.Type)
		}
	}
	return sections, nil
}

func (m *MessageSections) FromAnthropicMessageSections(sections []anthropic.ContentBlockUnion) error {
	for _, section := range sections {
		switch section.AsAny().(type) {
		case anthropic.TextBlock:
			*m = append(*m, MessageSection{
				Type:    MessageSectionTypeText,
				Payload: MessageTextSection{Text: section.Text},
			})
		case anthropic.RedactedThinkingBlock:
		case anthropic.ThinkingBlock:
		case anthropic.ToolUseBlock:
		case anthropic.ServerToolUseBlock:
		case anthropic.WebSearchToolResultBlock:
		default:
			return fmt.Errorf("unsupported anthropic message type")
		}
	}
	return nil
}

func (m *MessageSections) UnmarshalJSON(data []byte) error {
	type AuxSection struct {
		Type    MessageSectionType `json:"type"`
		Payload json.RawMessage    `json:"payload"`
	}
	var sections []AuxSection
	if err := json.Unmarshal(data, &sections); err != nil {
		return err
	}

	for _, section := range sections {
		var messageSection MessageSection
		switch section.Type {
		case MessageSectionTypeText:
			var textContent MessageTextSection
			if err := json.Unmarshal(section.Payload, &textContent); err != nil {
				return err
			}
			messageSection.Type = MessageSectionTypeText
			messageSection.Payload = textContent
		// TODO: handle more content types
		default:
			return fmt.Errorf("unknown message content type: %s", section.Type)
		}
		*m = append(*m, messageSection)
	}

	return nil
}

func (m *MessageSections) MarshalJSON() ([]byte, error) {
	type AuxSection struct {
		Type    MessageSectionType `json:"type"`
		Payload json.RawMessage    `json:"payload"`
	}
	var sections []AuxSection

	for _, section := range *m {
		var payload json.RawMessage
		switch section.Type {
		case MessageSectionTypeText:
			payloadRes, err := json.Marshal(section.Payload.(MessageTextSection))
			if err != nil {
				return nil, err
			}
			payload = payloadRes
		// TODO: handle more content types
		default:
			return nil, fmt.Errorf("unknown message content type: %s", section.Type)
		}
		sections = append(sections, AuxSection{
			Type:    section.Type,
			Payload: payload,
		})
	}
	return json.Marshal(sections)
}
