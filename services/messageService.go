package services

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

type messageService struct {
	messageRepo      domain.MessageRepository
	conversationRepo domain.ConversationRepository
}

func NewMessageService(messageRepo domain.MessageRepository, conversationRepo domain.ConversationRepository) domain.MessageService {
	return &messageService{messageRepo, conversationRepo}
}

func (m *messageService) Create(sender, convID uuid.UUID, content string) (domain.Message, error) {
	if ok, err := m.isUserInConversation(sender, convID); err != nil {
		return domain.Message{}, errors.New("conversation not found")
	} else if !ok {
		return domain.Message{}, errors.New("user not in conversation")
	}

	id, err := uuid.NewUUID()

	if err != nil {
		return domain.Message{}, errors.New("failed to generate message ID")
	}

	msg := domain.Message{
		ID:             id,
		SenderID:       sender,
		ConversationID: convID,
		Content:        content,
		CreatedAt:      time.Now(),
	}

	if err := m.messageRepo.Save(&msg); err != nil {
		return domain.Message{}, errors.New("failed to save message")
	}

	return msg, nil
}

func (m *messageService) GetAllFrom(user uuid.UUID, conversation uuid.UUID) ([]domain.Message, error) {
	if ok, err := m.isUserInConversation(user, conversation); err != nil || !ok {
		return nil, errors.New("failed to verify user in conversation") // Generic way xD
	}

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

	if oldMsg.SenderID == user {
		return errors.New("cannot react to own message")
	}

	if !oldMsg.DeletedAt.IsZero() {
		return errors.New("cannot react to deleted message")
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

	if !oldMsg.DeletedAt.IsZero() {
		return errors.New("cannot edit deleted message")
	}

	oldMsg.Content = newContent
	oldMsg.EditedAt = time.Now()

	if err := m.messageRepo.Update(msg, &oldMsg); err != nil {
		return errors.New("failed to edit message")
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

// Helpers

func (m *messageService) isUserInConversation(userID, convID uuid.UUID) (bool, error) {
	conv, err := m.conversationRepo.FindByID(convID)

	if err != nil {
		return false, errors.New("conversation not found")
	}

	return conv.UserID_A == userID || conv.UserID_B == userID, nil
}
