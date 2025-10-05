package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID             uuid.UUID `json:"id"`
	ConversationID uuid.UUID `json:"conversation_id"` // FK
	SenderID       uuid.UUID `json:"sender_id"`       // FK
	Content        string    `json:"content"`
	Reaction       string    `json:"reaction"`
	CreatedAt      time.Time `json:"created_at"`
	EditedAt       time.Time `json:"edited_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

type MessageService interface {
	Create(sender, conversation uuid.UUID, content string) error
	Edit(sender, msg uuid.UUID, newContent string) error
	React(user, msg uuid.UUID, reaction string) error
	Delete(sender, msg uuid.UUID) error
	GetAllFrom(user, conversation uuid.UUID) ([]Message, error)
}

type MessageRepository interface {
	Save(msg *Message) error
	FindAllFrom(conversation uuid.UUID) ([]Message, error)
	FindByID(id uuid.UUID) (Message, error)
	Update(id uuid.UUID, newMsg *Message) error
}
