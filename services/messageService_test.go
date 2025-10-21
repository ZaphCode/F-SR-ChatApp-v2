package services

import (
	"testing"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
	"github.com/ZaphCode/F-SR-ChatApp/utils/mocks"
)

type CreateMessageInput struct {
	SenderID       uuid.UUID
	ConversationID uuid.UUID
	Content        string
}

func TestMessageService_Create(t *testing.T) {
	tcs := []utils.TestCase[CreateMessageInput, domain.Message]{
		{
			Name: "Invalid sender ID",
			Input: CreateMessageInput{
				SenderID:       mocks.UserC.ID,
				ConversationID: mocks.ConversationUserAUserB.ID,
				Content:        "Hello, world!",
			},
			ExpectError: true,
		},
		{
			Name: "Invalid conversation ID",
			Input: CreateMessageInput{
				SenderID:       mocks.UserA.ID,
				ConversationID: uuid.New(),
				Content:        "Hello, world!",
			},
			ExpectError: true,
		},
		{
			Name: "Valid message",
			Input: CreateMessageInput{
				SenderID:       mocks.UserA.ID,
				ConversationID: mocks.ConversationUserAUserB.ID,
				Content:        "Hello, world!",
			},
			ExpectError: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			out, err := testMessageService.Create(
				tc.Input.SenderID,
				tc.Input.ConversationID,
				tc.Input.Content,
			)

			if (err != nil) != tc.ExpectError {
				t.Errorf("expected error: %v, got: %v", tc.ExpectError, err)
			}

			if err != nil {
				utils.PrettyPrint(err)
				return
			}

			utils.PrettyPrint(out)

			if tc.HandleOutput != nil {
				tc.HandleOutput(t, out)
			}
		})
	}
}

type GetAllInput struct {
	UserID         uuid.UUID
	ConversationID uuid.UUID
}

func TestMessageService_GetAllFrom(t *testing.T) {
	tcs := []utils.TestCase[GetAllInput, []domain.Message]{
		{
			Name: "Invalid user ID",
			Input: GetAllInput{
				UserID:         uuid.New(),
				ConversationID: mocks.ConversationUserAUserB.ID,
			},
			ExpectError: true,
		},
		{
			Name: "Invalid conversation ID",
			Input: GetAllInput{
				UserID:         mocks.UserA.ID,
				ConversationID: uuid.New(),
			},
			ExpectError: true,
		},
		{
			Name: "Valid user A messages",
			Input: GetAllInput{
				UserID:         mocks.UserA.ID,
				ConversationID: mocks.ConversationUserAUserB.ID,
			},
			ExpectError: false,
		},
		{
			Name: "Valid user B messages",
			Input: GetAllInput{
				UserID:         mocks.UserB.ID,
				ConversationID: mocks.ConversationUserAUserB.ID,
			},
			ExpectError: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			out, err := testMessageService.GetAllFrom(tc.Input.UserID, tc.Input.ConversationID)

			if (err != nil) != tc.ExpectError {
				t.Errorf("expected error: %v, got: %v", tc.ExpectError, err)
			}

			if err != nil {
				utils.PrettyPrint(err)
				return
			}

			utils.PrettyPrint(out)

			if tc.HandleOutput != nil {
				tc.HandleOutput(t, out)
			}
		})
	}
}
