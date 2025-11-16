package mappers

import (
	"arukabe/core/common/types"
	modelsv1 "arukabe/gen/connect/models/v1"
)

func PBMessageContentToMessageSections(content []*modelsv1.MessageContent) types.MessageSections {
	var sections types.MessageSections
	for _, item := range content {
		textContent := item.GetTextContent()
		if textContent != nil {
			sections = append(sections, types.MessageSection{
				Type:    types.MessageSectionTypeText,
				Payload: types.MessageTextSection{Text: textContent.Text},
			})
			continue
		}
		// TODO: handle more content types
	}

	return sections
}
