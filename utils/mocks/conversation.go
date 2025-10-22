package mocks

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

//* Conversations Mock

var ConversationUserAUserB domain.Conversation = domain.Conversation{
	ID:        uuid.MustParse("33333333-3333-3333-3333-333333333333"),
	UserA:     UserA,
	UserB:     UserB,
	CreatedAt: time.Now(),
}

//* Conversation Repository Mock

type conversationRepositoryMock struct {
	mu            sync.RWMutex
	conversations map[uuid.UUID]domain.Conversation
}

func NewConversationRepository() domain.ConversationRepository {
	return &conversationRepositoryMock{
		mu: sync.RWMutex{},
		conversations: map[uuid.UUID]domain.Conversation{
			ConversationUserAUserB.ID: ConversationUserAUserB,
		},
	}
}

func (c *conversationRepositoryMock) FindAllFrom(userID uuid.UUID) ([]domain.Conversation, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var result []domain.Conversation
	for _, conv := range c.conversations {
		if conv.UserA.ID == userID || conv.UserB.ID == userID {
			result = append(result, conv)
		}
	}
	return result, nil
}

func (c *conversationRepositoryMock) FindFrom(userAID uuid.UUID, userBID uuid.UUID) (domain.Conversation, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, conv := range c.conversations {
		if (conv.UserA.ID == userAID && conv.UserB.ID == userBID) ||
			(conv.UserA.ID == userBID && conv.UserB.ID == userAID) {
			return conv, nil
		}
	}
	return domain.Conversation{}, errors.New("conversation not found")
}

func (c *conversationRepositoryMock) Save(conv *domain.Conversation) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.conversations[conv.ID] = *conv
	return nil
}

func (c *conversationRepositoryMock) FindByID(id uuid.UUID) (domain.Conversation, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	conv, exists := c.conversations[id]
	if !exists {
		return domain.Conversation{}, errors.New("conversation not found")
	}
	return conv, nil
}
