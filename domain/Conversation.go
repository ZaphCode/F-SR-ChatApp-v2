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
