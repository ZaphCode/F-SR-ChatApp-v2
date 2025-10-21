package services

import (
	"testing"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
	"github.com/ZaphCode/F-SR-ChatApp/utils/mocks"
)

func TestUserService_Create(t *testing.T) {
	newUser, err := testUserService.Create("test-user", "testuser@example.com", "password123")

	if err != nil {
		t.Fatalf("unexpected error creating user: %v", err)
	}

	if newUser.ID == uuid.Nil {
		t.Error("expected valid UUID, got Nil")
	}

	if newUser.CreatedAt.IsZero() {
		t.Error("expected valid CreatedAt, got zero value")
	}

	if newUser.Password == "password123" {
		t.Error("expected hashed password, got plain text")
	}

	utils.PrettyPrint(newUser)

	_, err = testUserService.Create("JohnDoe", newUser.Email, "password456")

	if err == nil {
		t.Error("expected error when creating user with existing email, got nil")
	}

	utils.PrettyPrint(err)
}

func TestUserService_GetByID(t *testing.T) {
	if _, err := testUserService.GetByID(uuid.New()); err == nil {
		t.Error("expected error for non-existent user ID, got none")
	}

	user, err := testUserService.GetByID(mocks.UserA.ID)

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	if user.Email != mocks.UserA.Email {
		t.Errorf("expected email: %v, got: %v", mocks.UserA.Email, user.Email)
	}

	utils.PrettyPrint(user)
}

type AuthenticateInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestUserService_Authenticate(t *testing.T) {
	tcs := []utils.TestCase[AuthenticateInput, domain.User]{
		{
			Name: "Authenticate with unregistered email",
			Input: AuthenticateInput{
				Email:    mocks.UserA.Email,
				Password: "xxxxxx",
			},
			ExpectError: true,
		},
		{
			Name: "Authenticate with invalid password",
			Input: AuthenticateInput{
				Email:    mocks.UserA.Email,
				Password: "wrongpassword",
			},
			ExpectError: true,
		},
		{
			Name: "Authenticate with valid credentials",
			Input: AuthenticateInput{
				Email:    mocks.UserA.Email,
				Password: "password123",
			},
			ExpectError: false,
			HandleOutput: func(t *testing.T, output domain.User) {
				if output.ID != mocks.UserA.ID {
					t.Errorf("expected ID: %v, got: %v", mocks.UserA.ID, output.ID)
				}

				utils.PrettyPrint(output)
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			out, err := testUserService.Authenticate(tc.Input.Email, tc.Input.Password)

			if (err != nil) != tc.ExpectError {
				t.Errorf("expected error: %v, got: %v", tc.ExpectError, err)
			}

			utils.PrettyPrint(err)

			if tc.HandleOutput != nil {
				tc.HandleOutput(t, out)
			}
		})
	}
}

func TestUserService_GetAll(t *testing.T) {
	users, err := testUserService.GetAll()

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	if len(users) == 0 {
		t.Error("expected at least one user, got zero")
	}

	utils.PrettyPrint(users)
}

func TestUserService_UpdateProfileImg(t *testing.T) {
	if err := testUserService.UpdateProfileImg(uuid.New(), "image.png"); err == nil {
		t.Fatalf("expected error for non-existent user ID, got none")
	}

	user, err := testUserService.GetByID(mocks.UserA.ID)

	if err != nil {
		t.Fatalf("unexpected error fetching user: %v", err)
	}

	utils.PrettyPrint(user)

	if err := testUserService.UpdateProfileImg(user.ID, "image.png"); err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	newUser, err := testUserService.GetByID(user.ID)

	if err != nil {
		t.Fatalf("unexpected error fetching user: %v", err)
	}

	if newUser.ImageUrl != "image.png" {
		t.Errorf("expected image URL to be updated to 'image.png', got: %v", newUser.ImageUrl)
	}

	utils.PrettyPrint(newUser)
}

func TestUserService_Delete(t *testing.T) {
	if err := testUserService.Delete(uuid.New()); err == nil {
		t.Fatalf("expected error for non-existent user ID, got none")
	}

	if err := testUserService.Delete(mocks.UserB.ID); err != nil {
		t.Errorf("expected no error deleting existing user, got: %v", err)
	}

	if _, err := testUserService.GetByID(mocks.UserB.ID); err == nil {
		t.Error("expected error fetching deleted user, got none")
	}
}
