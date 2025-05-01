package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID             uuid.UUID
	ConversationID uuid.UUID
	SenderID       uuid.UUID
	Content        string
	File           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type MessageService interface {
	Create(sender, conv uuid.UUIDs, content, file string) error
	GetFrom(conv uuid.UUID)
}

type MessageRepository interface {
	Save(msg *Message) error
	FindAllFrom(conv uuid.Domain) error
}
