package models

import "github.com/google/uuid"

type Chat struct {
	ID       uuid.UUID
	Title    string
	Prompt   string
	Pinned   bool
	Model    Model
	TimeData TimeData
	Messages []Message
}
