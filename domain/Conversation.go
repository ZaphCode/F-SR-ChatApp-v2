package domain

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	ID        uuid.UUID
	UserID_A  uuid.UUID
	UserID_B  uuid.UUID
	CreatedAt time.Time
}

type ConversationService interface {
	GetOrCreateFrom(userA, userB uuid.UUID) (Conversation, error)
	GetAllFrom(user uuid.UUID) ([]Conversation, error)
}

type ConversationRepository interface {
	Save(conv *Conversation) error
	FindFrom(userA, userB uuid.UUID) (Conversation, error)
	FindAllFrom(user_id uuid.UUID) ([]Conversation, error)
}
