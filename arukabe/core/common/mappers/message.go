package mappers

import (
	"arukabe/core/common/types"
	modelsv1 "arukabe/gen/connect/aruka/models/v1"
	"arukabe/gen/sqlc"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromDBMessage(message sqlc.Message) (types.Message, error) {
	var content types.MessageSections
	if err := content.UnmarshalJSON(message.Content); err != nil {
		return types.Message{}, err
	}
	return types.Message{
		ID:        message.ID,
		ChatID:    message.ChatID,
		Role:      message.Role,
		Content:   content,
		CreatedAt: message.CreatedAt,
	}, nil
}

func FromDBMessages(messages []sqlc.Message) ([]types.Message, error) {
	result := make([]types.Message, len(messages))
	for i := range messages {
		msg, err := FromDBMessage(messages[i])
		if err != nil {
			return nil, err
		}
		result[i] = msg
	}
	return result, nil
}

func MessageToPB(message types.Message) (*modelsv1.Message, error) {
	role, err := RoleToPB(message.Role)
	if err != nil {
		return nil, err
	}
	sections, err := message.Content.ToPBSections()
	if err != nil {
		return nil, err
	}
	protoMsg := &modelsv1.Message{
		Id:        message.ID.String(),
		Role:      role,
		Content:   sections,
		CreatedAt: timestamppb.New(message.CreatedAt),
	}
	return protoMsg, nil
}

func MessagesToPB(messages []types.Message) ([]*modelsv1.Message, error) {
	result := make([]*modelsv1.Message, len(messages))
	for i := range messages {
		pbMsg, err := MessageToPB(messages[i])
		if err != nil {
			return nil, err
		}
		result[i] = pbMsg
	}
	return result, nil
}

func DBMessagesToPB(messages []sqlc.Message) ([]*modelsv1.Message, error) {
	result := make([]*modelsv1.Message, len(messages))
	for i := range messages {
		msg, err := FromDBMessage(messages[i])
		if err != nil {
			return nil, err
		}
		pbMsg, err := MessageToPB(msg)
		if err != nil {
			return nil, err
		}
		result[i] = pbMsg
	}
	return result, nil
}
