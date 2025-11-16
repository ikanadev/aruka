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
