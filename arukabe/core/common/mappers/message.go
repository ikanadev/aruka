package mappers

import (
	"arukabe/core/common/types"
	"arukabe/gen/sqlc"
)

func FromDBMessages(messages []sqlc.Message) ([]types.Message, error) {
	result := make([]types.Message, len(messages))
	for i := range messages {
		if err := result[i].FromDBMessage(messages[i]); err != nil {
			return nil, err
		}
	}
	return result, nil
}
