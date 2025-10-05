package repositories

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

func TestMongoDBMessageRepository_Save(t *testing.T) {
	msg := domain.Message{
		ID:             uuid.New(),
		ConversationID: uuid.New(),
		SenderID:       uuid.New(),
		Content:        "Hello, World!",
		CreatedAt:      time.Now(),
	}

	if err := messageRepo.Save(&msg); err != nil {
		t.Fatalf("Failed to save message: %v", err)
	}

	utils.PrettyPrint(msg)
}

func TestMongoDBMessageRepository_SaveAndFindAllFrom(t *testing.T) {
	convID := uuid.New()

	for i := 0; i < 2; i++ {
		if err := messageRepo.Save(&domain.Message{
			ID:             uuid.New(),
			ConversationID: convID,
			SenderID:       uuid.New(),
			Content:        fmt.Sprintf("Message %d", i+1),
			CreatedAt:      time.Now(),
		}); err != nil {
			t.Fatalf("Failed to save message: %v", err)
		}
	}

	messages, err := messageRepo.FindAllFrom(convID)

	if err != nil {
		t.Fatalf("Failed to find messages: %v", err)
	}

	if len(messages) != 2 {
		t.Fatalf("Expected 2 messages, got %d", len(messages))
	}

	utils.PrettyPrint(messages)
}

func TestMongoDBMessageRepository_SaveAndUpdate(t *testing.T) {
	msg := domain.Message{
		ID:             uuid.New(),
		ConversationID: uuid.New(),
		SenderID:       uuid.New(),
		Content:        "Original Content",
		CreatedAt:      time.Now(),
	}

	if err := messageRepo.Save(&msg); err != nil {
		t.Fatalf("Failed to save message: %v", err)
	}

	msg.Content = "Updated Content"
	msg.Reaction = "👍"
	msg.EditedAt = time.Now()

	if err := messageRepo.Update(msg.ID, &msg); err != nil {
		t.Fatalf("Failed to update message: %v", err)
	}

	utils.PrettyPrint(msg)

}
