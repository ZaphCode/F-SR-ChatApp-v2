package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID             uuid.UUID `json:"id"`
	ConversationID uuid.UUID `json:"conversation_id"`
	SenderID       uuid.UUID `json:"sender_id"`
	Content        string    `json:"content"`
	File           string    `json:"file"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MessageService interface {
	Create(sender, conv uuid.UUIDs, content, file string) error
	GetFrom(conv uuid.UUID)
}

type MessageRepository interface {
	Save(msg *Message) error
	FindAllFrom(conv uuid.Domain) error
}
