package mappers

import (
	modelsv1 "arukabe/gen/connect/models/v1"
	"arukabe/gen/sqlc"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

func RoleToAnthropicRoleParam(role sqlc.MessageRole) (anthropic.MessageParamRole, error) {
	switch role {
	case sqlc.MessageRoleASSISTANT:
		return anthropic.MessageParamRoleAssistant, nil
	case sqlc.MessageRoleUSER:
		return anthropic.MessageParamRoleUser, nil
	default:
		return "", fmt.Errorf("unknown role: %s", role)
	}
}

func RoleToPB(role sqlc.MessageRole) (modelsv1.MessageRole, error) {
	switch role {
	case sqlc.MessageRoleASSISTANT:
		return modelsv1.MessageRole_MESSAGE_ROLE_ASSISTANT, nil
	case sqlc.MessageRoleUSER:
		return modelsv1.MessageRole_MESSAGE_ROLE_USER, nil
	default:
		return modelsv1.MessageRole_MESSAGE_ROLE_UNSPECIFIED, fmt.Errorf("unknown role: %s", role)
	}
}
