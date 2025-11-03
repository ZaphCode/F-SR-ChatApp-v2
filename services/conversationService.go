package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

type conversationService struct {
	conversationRepo domain.ConversationRepository
	userRepository   domain.UserRepository
}

func NewConversationService(
	conversationRepo domain.ConversationRepository,
	userRepository domain.UserRepository,
) domain.ConversationService {
	return &conversationService{conversationRepo, userRepository}
}

func (s *conversationService) GetOrCreateFrom(userAID, userBID uuid.UUID) (domain.Conversation, error) {
	if userAID == userBID {
		return domain.Conversation{}, fmt.Errorf("userA and userB are the same")
	}

	conv, err := s.conversationRepo.FindFrom(userAID, userBID)
	if err == nil {
		return conv, nil
	}

	errChan := make(chan error, 2)

	go func() {
		conv.UserA, err = s.userRepository.FindByID(userAID)
		errChan <- err
	}()

	go func() {
		conv.UserB, err = s.userRepository.FindByID(userBID)
		errChan <- err
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

func (s *conversationService) GetAllFrom(userID uuid.UUID) ([]domain.Conversation, error) {
	conversations, err := s.conversationRepo.FindAllFrom(userID)

	if err != nil {
		return nil, err
	}

	cache := sync.Map{}
	wg := sync.WaitGroup{}

	for i := range conversations {
		wg.Add(2)

		go func(conv *domain.Conversation) {
			defer wg.Done()
			if user, err := s.getUserWithCaching(&cache, conv.UserA.ID); err == nil {
				conv.UserA = user
			}
		}(&conversations[i])

		go func(conv *domain.Conversation) {
			defer wg.Done()
			if user, err := s.getUserWithCaching(&cache, conv.UserB.ID); err == nil {
				conv.UserB = user
			}
		}(&conversations[i])
	}

	wg.Wait()

	return conversations, nil
}

func (s *conversationService) getUserWithCaching(cache *sync.Map, userID uuid.UUID) (domain.User, error) {
	if cached, ok := cache.Load(userID); ok {
		return cached.(domain.User), nil
	}

	user, err := s.userRepository.FindByID(userID)

	if err != nil {
		return domain.User{}, err
	}

	cache.Store(userID, user)

	return user, nil
}
