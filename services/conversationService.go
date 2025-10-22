package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

func NewConversationService(
	conversationRepo domain.ConversationRepository,
	userRepository domain.UserRepository,
) domain.ConversationService {
	return &conversationService{conversationRepo, userRepository}
}

type conversationService struct {
	conversationRepo domain.ConversationRepository
	userRepository   domain.UserRepository
}

func (s *conversationService) GetOrCreateFrom(userAID, userBID uuid.UUID) (domain.Conversation, error) {
	if userAID == userBID {
		return domain.Conversation{}, fmt.Errorf("userA and userB are the same")
	}

	var conv domain.Conversation

	conv, err := s.conversationRepo.FindFrom(userAID, userBID)

	if err != nil {
		errChan := make(chan error, 2)

		go func() {
			userA, err := s.userRepository.FindByID(userAID)
			errChan <- err
			conv.UserA = userA
		}()

		go func() {
			userB, err := s.userRepository.FindByID(userBID)
			errChan <- err
			conv.UserB = userB
		}()

		for range 2 {
			if err := <-errChan; err != nil {
				return domain.Conversation{}, fmt.Errorf("one or both users not found")
			}
		}

		id, err := uuid.NewUUID()

		if err != nil {
			return conv, err
		}

		conv.ID = id
		conv.CreatedAt = time.Now()

		if err := s.conversationRepo.Save(&conv); err != nil {
			return conv, err
		}

		return conv, nil
	}

	return conv, nil
}

func (s *conversationService) GetAllFrom(userID uuid.UUID) ([]domain.Conversation, error) {
	return s.conversationRepo.FindAllFrom(userID)
}
