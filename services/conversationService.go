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

	conv, err := s.conversationRepo.FindFrom(userAID, userBID)

	if err != nil {
		errChan := make(chan error, 2)

		check := func(id uuid.UUID) {
			_, err := s.userRepository.FindByID(id)
			errChan <- err
		}

		go check(userAID)
		go check(userBID)

		for range 2 {
			if err := <-errChan; err != nil {
				return domain.Conversation{}, fmt.Errorf("one or both users not found")
			}
		}

		id, err := uuid.NewUUID()

		if err != nil {
			return conv, err
		}

		newConv := domain.Conversation{
			ID:        id,
			UserID_A:  userAID,
			UserID_B:  userBID,
			CreatedAt: time.Now(),
		}

		if err := s.conversationRepo.Save(&newConv); err != nil {
			return conv, err
		}

		return newConv, nil
	}

	return conv, nil
}

func (s *conversationService) GetAllFrom(userID uuid.UUID) ([]domain.Conversation, error) {
	return s.conversationRepo.FindAllFrom(userID)
}
