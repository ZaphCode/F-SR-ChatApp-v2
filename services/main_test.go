package services

import (
	"testing"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/ZaphCode/F-SR-ChatApp/utils/mocks"
)

// var testConversationService domain.ConversationService
// var testMessageService domain.MessageService
var testUserService domain.UserService

func TestMain(m *testing.M) {
	testUserService = NewUserService(mocks.NewUserRepository())

	m.Run()
}
