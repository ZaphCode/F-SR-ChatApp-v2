package handlers

import (
	"net/http"
	"testing"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/ZaphCode/F-SR-ChatApp/services"
	"github.com/ZaphCode/F-SR-ChatApp/utils/mocks"
)

var testingMux *http.ServeMux

func TestMain(m *testing.M) {
	app.InitSessionStore()
	testingMux = http.NewServeMux()

	// Repositories
	userRepository := mocks.NewUserRepository()

	// Services
	userService := services.NewUserService(userRepository)

	NewAuthHandler(userService).SetRoutes(testingMux)

	m.Run()
}
