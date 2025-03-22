package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID             uuid.UUID
	Content        string
	File           string
	ConversationID string
	SenderID       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
