package repo

import (
	"arukabe/domain/models"

	"github.com/google/uuid"
)

type NewChatData struct {
	Prompt  string
	ModelID uuid.UUID
}
type ChatRepository interface {
  NewChat(data NewChatData) (models.Chat, error)
}
