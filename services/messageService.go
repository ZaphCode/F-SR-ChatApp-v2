package services

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

type messageService struct {
	messageRepo domain.MessageRepository
}

func NewMessageService(messageRepo domain.MessageRepository) domain.MessageService {
	return &messageService{messageRepo}
}

func (m *messageService) Create(sender, conv uuid.UUID, content string) error {
	msg := domain.Message{
		ID:             uuid.New(),
		SenderID:       sender,
		ConversationID: conv,
		Content:        content,
		CreatedAt:      time.Now(),
	}

	if err := m.messageRepo.Save(&msg); err != nil {
		return errors.New("failed to save message")
	}

	return nil
}

func (m *messageService) Delete(sender uuid.UUID, msg uuid.UUID) error {
	oldMsg, err := m.messageRepo.FindByID(msg)

	if err != nil {
		return errors.New("message not found")
	}

	if oldMsg.SenderID != sender {
		return errors.New("not authorized")
	}

	oldMsg.DeletedAt = time.Now()

	if err := m.messageRepo.Update(msg, &oldMsg); err != nil {
		return errors.New("failed to delete message")
	}

	return nil
}

func (m *messageService) GetAllFrom(user uuid.UUID, conversation uuid.UUID) ([]domain.Message, error) {
	msgs, err := m.messageRepo.FindAllFrom(conversation)

	if err != nil {
		return nil, errors.New("failed to get messages")
	}

	return msgs, nil
}

func (m *messageService) React(user uuid.UUID, msg uuid.UUID, reaction string) error {
	oldMsg, err := m.messageRepo.FindByID(msg)

	if err != nil {
		return errors.New("message not found")
	}

	oldMsg.Reaction = reaction

	if err := m.messageRepo.Update(msg, &oldMsg); err != nil {
		return errors.New("failed to react to message")
	}

	return nil
}

func (m *messageService) Edit(sender uuid.UUID, msg uuid.UUID, newContent string) error {
	oldMsg, err := m.messageRepo.FindByID(msg)

	if err != nil {
		return errors.New("message not found")
	}

	if oldMsg.SenderID != sender {
		return errors.New("not authorized")
	}

	oldMsg.Content = newContent
	oldMsg.EditedAt = time.Now()

	if err := m.messageRepo.Update(msg, &oldMsg); err != nil {
		return errors.New("failed to edit message")
	}

	return nil
}
