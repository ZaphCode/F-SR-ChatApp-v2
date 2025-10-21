package mocks

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

var MessageUserAtoUserB_1 domain.Message = domain.Message{
	ID:             uuid.MustParse("44444444-4444-4444-4444-444444444444"),
	ConversationID: ConversationUserAUserB.ID,
	SenderID:       UserA.ID,
	Content:        "Hello, User B!",
	CreatedAt:      time.Now(),
}

var MessageUserBtoUserA_1 domain.Message = domain.Message{
	ID:             uuid.MustParse("55555555-5555-5555-5555-555555555555"),
	ConversationID: ConversationUserAUserB.ID,
	SenderID:       UserB.ID,
	Content:        "Hi, User A! How are you?",
	CreatedAt:      time.Now(),
}

var MessageUserAtoUserB_2 domain.Message = domain.Message{
	ID:             uuid.MustParse("66666666-6666-6666-6666-666666666666"),
	ConversationID: ConversationUserAUserB.ID,
	SenderID:       UserA.ID,
	Content:        "I'm doing well, thanks for asking!",
	CreatedAt:      time.Now(),
}

type messageRepositoryMock struct {
	mu       sync.RWMutex
	messages map[uuid.UUID]domain.Message
}

func NewMessageRepository() domain.MessageRepository {
	return &messageRepositoryMock{
		mu: sync.RWMutex{},
		messages: map[uuid.UUID]domain.Message{
			MessageUserAtoUserB_1.ID: MessageUserAtoUserB_1,
			MessageUserBtoUserA_1.ID: MessageUserBtoUserA_1,
			MessageUserAtoUserB_2.ID: MessageUserAtoUserB_2,
		},
	}
}

func (m *messageRepositoryMock) FindAllFrom(conversation uuid.UUID) ([]domain.Message, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var messages []domain.Message
	for _, msg := range m.messages {
		if msg.ConversationID == conversation {
			messages = append(messages, msg)
		}
	}
	return messages, nil
}

func (m *messageRepositoryMock) FindByID(id uuid.UUID) (domain.Message, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	msg, exists := m.messages[id]
	if !exists {
		return domain.Message{}, errors.New("message not found")
	}
	return msg, nil
}

func (m *messageRepositoryMock) Save(msg *domain.Message) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.messages[msg.ID] = *msg
	return nil
}

func (m *messageRepositoryMock) Update(id uuid.UUID, newMsg *domain.Message) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exists := m.messages[id]

	if !exists {
		return errors.New("message not found")
	}

	m.messages[id] = *newMsg

	return nil
}
