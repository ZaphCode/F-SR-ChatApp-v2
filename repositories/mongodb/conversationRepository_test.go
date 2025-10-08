package mongodb

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

func TestMongoDBConversationRepository_Save(t *testing.T) {
	conv := &domain.Conversation{
		ID:        uuid.New(),
		UserID_A:  uuid.New(),
		UserID_B:  uuid.New(),
		CreatedAt: time.Now(),
	}

	err := conversationRepo.Save(conv)

	if err != nil {
		t.Fatalf("Failed to save conversation: %v", err)
	}
}

func TestMongoDBConversationRepository_SaveAndFindFrom(t *testing.T) {
	conv := &domain.Conversation{
		ID:        uuid.New(),
		UserID_A:  uuid.New(),
		UserID_B:  uuid.New(),
		CreatedAt: time.Now(),
	}

	err := conversationRepo.Save(conv)

	if err != nil {
		t.Fatalf("Failed to save conversation: %v", err)
	}

	result, err := conversationRepo.FindFrom(conv.UserID_A, conv.UserID_B)

	if err != nil {
		t.Fatalf("Failed to find conversation: %v", err)
	}

	if result.ID != conv.ID {
		t.Fatalf("Expected conversation ID %s, got %s", conv.ID, result.ID)
	}

	result, err = conversationRepo.FindFrom(conv.UserID_B, conv.UserID_A)

	if err != nil {
		t.Fatalf("Failed to find conversation: %v", err)
	}

	if result.ID != conv.ID {
		t.Fatalf("Expected conversation ID %s, got %s", conv.ID, result.ID)
	}
}

func TestMongoDBConversationRepository_SaveAndFindAll(t *testing.T) {
	primaryUserID := uuid.New()

	for range 2 {
		err := conversationRepo.Save(&domain.Conversation{
			ID:        uuid.New(),
			UserID_A:  primaryUserID,
			UserID_B:  uuid.New(),
			CreatedAt: time.Now(),
		})

		if err != nil {
			t.Fatalf("Failed to save conversation: %v", err)
		}
	}

	for range 2 {
		err := conversationRepo.Save(&domain.Conversation{
			ID:        uuid.New(),
			UserID_A:  uuid.New(),
			UserID_B:  primaryUserID,
			CreatedAt: time.Now(),
		})

		if err != nil {
			t.Fatalf("Failed to save conversation: %v", err)
		}
	}

	conversations, err := conversationRepo.FindAllFrom(primaryUserID)

	if err != nil {
		t.Fatalf("Failed to find all conversations: %v", err)
	}

	if len(conversations) != 4 {
		t.Fatalf("Expected 4 conversations, got %d", len(conversations))
	}

	for _, c := range conversations {
		t.Logf("Conversation ID: %s, User A: %s, User B: %s", c.ID, c.UserID_A, c.UserID_B)
	}
}

// Sample UUIDs for testing
// d8900901-5efc-4b5d-b7cb-a9b8910bc3d7 user a 1
// 102b6459-e829-429f-8357-2f8b8bb60367 user b 1
// 7f973f86-8900-49eb-8ea9-a5b4be0aa55c user a 2
// 6248f644-9116-4fad-99c6-559da36f3fea user b 2
// d56d642c-af83-4942-ac93-12f747c84926 EXTRA
