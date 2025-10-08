package services

import (
	"testing"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestUserService_Create(t *testing.T) {
	tcs := []utils.TestCase[CreateUserInput, domain.User]{
		{
			Name: "Create User Successfully",
			Input: CreateUserInput{
				Username: "testuser",
				Email:    "testuser@example.com",
				Password: "password123",
			},
			ExpectError: false,
			HandleOutput: func(t *testing.T, output domain.User) {
				if output.ID == uuid.Nil {
					t.Error("expected valid UUID, got Nil")
				}

				if output.CreatedAt.IsZero() {
					t.Error("expected valid CreatedAt, got zero value")
				}

				if output.Password == "password123" {
					t.Error("expected hashed password, got plain text")
				}

				utils.PrettyPrint(output)
			},
		},
		{
			Name: "Create User With Existing Email",
			Input: CreateUserInput{
				Username: "JohnDoe",
				Email:    "testuser@example.com",
				Password: "password456",
			},
			ExpectError: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			out, err := testUserService.Create(tc.Input.Username, tc.Input.Email, tc.Input.Password)

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
