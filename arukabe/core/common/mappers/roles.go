package mappers

import (
	"arukabe/gen/sqlc"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

func RoleToAnthropic(role sqlc.MessageRole) (anthropic.MessageParamRole, error) {
	switch role {
	case sqlc.MessageRoleASSISTANT:
		return anthropic.MessageParamRoleAssistant, nil
	case sqlc.MessageRoleUSER:
		return anthropic.MessageParamRoleUser, nil
	default:
		return "", fmt.Errorf("unknown role: %s", role)
	}
}
