package models

import (
	"time"

	"github.com/google/uuid"
)

type MessageRole string

const (
	MessageRoleUser      MessageRole = "USER"
	MessageRoleAssistant MessageRole = "ASSISTANT"
)

type MessageType string

const (
	MessageTypeText MessageType = "TEXT"
)

type Message struct {
	ID        uuid.UUID
	Role      MessageRole
	CreatedAt time.Time
}

type MessageContent interface {
	Type() MessageType
}

// Message content types
type TextMessageContent struct {
	Text string
}

func (t TextMessageContent) Type() MessageType {
	return MessageTypeText
}
