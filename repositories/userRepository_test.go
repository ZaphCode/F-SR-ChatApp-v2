package repositories

import (
	"errors"
	"testing"
	"time"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
	"github.com/google/uuid"
)

func Test_mongoDBUserRepo_FindByEmail(t *testing.T) {
	user := domain.User{
		ID:        uuid.New(),
		Username:  "Test User",
		Email:     "omar@gmail.com",
		CreatedAt: time.Now(),
		Password:  "password",
	}

	err := userRepo.Save(&user)

	if err != nil {
		t.Fatalf("Failed to save user: %v", err)
	}

	foundUser, err := userRepo.FindByEmail(user.Email)

	if err != nil {
		t.Errorf("Failed to find user by email: %v", err)
	}

	if foundUser.ID != user.ID {
		t.Errorf("Expected user ID %v, got %v", user.ID, foundUser.ID)
	}

	if _, err := userRepo.FindByEmail("random@gmail.mx"); !errors.Is(err, utils.ErrNotFound) {
		t.Errorf("Expected not found error, got %v", err)
	}
}
