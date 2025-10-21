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

type ReactMessageInput struct {
	UserID    uuid.UUID
	MessageID uuid.UUID
	Reaction  string
}

func TestMessageService_React(t *testing.T) {
	tcs := []utils.TestCase[ReactMessageInput, domain.Message]{
		{
			Name: "React to non-existent message",
			Input: ReactMessageInput{
				UserID:    mocks.UserA.ID,
				MessageID: uuid.New(),
				Reaction:  "👍",
			},
			ExpectError: true,
		},
		{
			Name: "React to own message",
			Input: ReactMessageInput{
				UserID:    mocks.UserA.ID,
				MessageID: mocks.MessageUserAtoUserB_1.ID,
				Reaction:  "❤️",
			},
			ExpectError: true,
		},
		{
			Name: "React to deleted message",
			Input: ReactMessageInput{
				UserID:    mocks.UserA.ID,
				MessageID: mocks.MessageUserBtoUserA_Deleted.ID,
				Reaction:  "😢",
			},
			ExpectError: true,
		},
		{
			Name: "Valid reaction",
			Input: ReactMessageInput{
				UserID:    mocks.UserB.ID,
				MessageID: mocks.MessageUserAtoUserB_1.ID,
				Reaction:  "😂",
			},
			ExpectError: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			err := testMessageService.React(tc.Input.UserID, tc.Input.MessageID, tc.Input.Reaction)

			if (err != nil) != tc.ExpectError {
				t.Errorf("expected error: %v, got: %v", tc.ExpectError, err)
			}

			if err != nil {
				utils.PrettyPrint(err)
				return
			}

			allMsgs, err := testMessageService.GetAllFrom(mocks.UserB.ID, mocks.ConversationUserAUserB.ID)

			if err != nil {
				t.Errorf("failed to get messages after reaction: %v", err)
				return
			}

			utils.PrettyPrint(allMsgs)
		})
	}
}

type EditMessageInput struct {
	SenderID   uuid.UUID
	MessageID  uuid.UUID
	NewContent string
}

func TestMessageService_Edit(t *testing.T) {
	tcs := []utils.TestCase[EditMessageInput, domain.Message]{
		{
			Name: "Edit non-existent message",
			Input: EditMessageInput{
				SenderID:   mocks.UserA.ID,
				MessageID:  uuid.New(),
				NewContent: "Edited content",
			},
			ExpectError: true,
		},
		{
			Name: "Edit message by non-sender",
			Input: EditMessageInput{
				SenderID:   mocks.UserB.ID,
				MessageID:  mocks.MessageUserAtoUserB_1.ID,
				NewContent: "Edited content",
			},
			ExpectError: true,
		},
		{
			Name: "Edit deleted message",
			Input: EditMessageInput{
				SenderID:   mocks.UserB.ID,
				MessageID:  mocks.MessageUserBtoUserA_Deleted.ID,
				NewContent: "Edited content",
			},
			ExpectError: true,
		},
		{
			Name: "Valid edit",
			Input: EditMessageInput{
				SenderID:   mocks.UserA.ID,
				MessageID:  mocks.MessageUserAtoUserB_2.ID,
				NewContent: "Edited content",
			},
			ExpectError: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			err := testMessageService.Edit(tc.Input.SenderID, tc.Input.MessageID, tc.Input.NewContent)

			if (err != nil) != tc.ExpectError {
				t.Errorf("expected error: %v, got: %v", tc.ExpectError, err)
			}

			if err != nil {
				utils.PrettyPrint(err)
				return
			}

			allMsgs, err := testMessageService.GetAllFrom(mocks.UserA.ID, mocks.ConversationUserAUserB.ID)

			if err != nil {
				t.Errorf("failed to get messages after edit: %v", err)
				return
			}

			utils.PrettyPrint(allMsgs)
		})
	}
}

func TestMessageService_Delete(t *testing.T) {
	err := testMessageService.Delete(mocks.UserA.ID, mocks.MessageUserBtoUserA_1.ID)

	if err == nil {
		t.Fatalf("should be error")
	}

	err = testMessageService.Delete(mocks.UserA.ID, mocks.MessageUserAtoUserB_2.ID)

	if err != nil {
		t.Fatalf("failed to delete message: %v", err)
	}

	allMsgs, err := testMessageService.GetAllFrom(mocks.UserA.ID, mocks.ConversationUserAUserB.ID)

	if err != nil {
		t.Fatalf("failed to get messages after delete: %v", err)
	}

	utils.PrettyPrint(allMsgs)
}
