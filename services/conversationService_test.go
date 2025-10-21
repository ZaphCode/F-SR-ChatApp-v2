package services

import (
	"testing"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
	"github.com/ZaphCode/F-SR-ChatApp/utils/mocks"
)

type GetOrCreateConversationInput struct {
	UserAID uuid.UUID
	UserBID uuid.UUID
}

func TestConversationService_GetOrCreateFrom(t *testing.T) {
	tcs := []utils.TestCase[GetOrCreateConversationInput, domain.Conversation]{
		{
			Name: "Create conversation with same user IDs",
			Input: GetOrCreateConversationInput{
				UserAID: mocks.UserA.ID,
				UserBID: mocks.UserA.ID,
			},
			ExpectError: true,
		},
		{
			Name: "Create conversation with non-existent user",
			Input: GetOrCreateConversationInput{
				UserAID: mocks.UserA.ID,
				UserBID: uuid.New(),
			},
			ExpectError: true,
		},
		{
			Name: "Get existing conversation between two different users",
			Input: GetOrCreateConversationInput{
				UserAID: mocks.UserA.ID,
				UserBID: mocks.UserB.ID,
			},
			ExpectError: false,
		},
		{
			Name: "Get existing conversation between two different users with reversed IDs",
			Input: GetOrCreateConversationInput{
				UserAID: mocks.UserB.ID,
				UserBID: mocks.UserA.ID,
			},
			ExpectError: false,
		},
		{
			Name: "Create new conversation between two different users",
			Input: GetOrCreateConversationInput{
				UserAID: mocks.UserA.ID,
				UserBID: mocks.UserC.ID,
			},
			ExpectError: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			output, err := testConversationService.GetOrCreateFrom(tc.Input.UserAID, tc.Input.UserBID)

			if tc.ExpectError {
				if err == nil {
					t.Error("expected error, got none")
				}
				return
			}

			if err != nil {
				t.Errorf("expected no error, got: %v", err)
				return
			}

			if (output.UserID_A != tc.Input.UserAID && output.UserID_A != tc.Input.UserBID) ||
				(output.UserID_B != tc.Input.UserAID && output.UserID_B != tc.Input.UserBID) {
				t.Errorf("conversation users do not match input IDs")
			}

			utils.PrettyPrint(output)
		})
	}
}

func TestConversationService_GetAllFrom(t *testing.T) {
	cs, err := testConversationService.GetAllFrom(uuid.New())

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(cs) != 0 {
		t.Fatalf("expected no conversations for unknown user, got: %v", cs)
	}

	cs, err = testConversationService.GetAllFrom(mocks.UserA.ID)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(cs) == 0 {
		t.Fatalf("expected at least one conversation, got none")
	}

	for _, conv := range cs {
		if conv.UserID_A != mocks.UserA.ID && conv.UserID_B != mocks.UserA.ID {
			t.Errorf("conversation does not belong to user A")
		}
	}

	utils.PrettyPrint(cs)
}
