package types

import (
	"arukabe/gen/sqlc"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID
	ChatID    uuid.UUID
	Role      sqlc.MessageRole
	Content   MessageSections
	CreatedAt time.Time
}

func (m *Message) FromDBMessage(messages sqlc.Message) error {
	var content MessageSections
	if err := content.UnmarshalJSON(messages.Content); err != nil {
		return err
	}
	m.ID = messages.ID
	m.ChatID = messages.ChatID
	m.Role = messages.Role
	m.Content = content
	m.CreatedAt = messages.CreatedAt
	return nil
}
