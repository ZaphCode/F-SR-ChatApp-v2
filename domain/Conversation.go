package domain

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	ID        uuid.UUID `json:"id"`
	UserID_A  uuid.UUID `json:"user_id_a"` // FK
	UserID_B  uuid.UUID `json:"user_id_b"` // FK
	CreatedAt time.Time `json:"created_at"`
}

type ConversationService interface {
	GetOrCreateFrom(userA, userB uuid.UUID) (Conversation, error)
	GetAllFrom(user uuid.UUID) ([]Conversation, error)
}

type ConversationRepository interface {
	Save(conv *Conversation) error
	FindFrom(userA, userB uuid.UUID) (Conversation, error)
	FindAllFrom(user uuid.UUID) ([]Conversation, error)
}
