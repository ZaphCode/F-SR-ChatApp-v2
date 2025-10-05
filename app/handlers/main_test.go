package handlers

import (
	"net/http"
	"testing"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/ZaphCode/F-SR-ChatApp/services"
)

var testingMux *http.ServeMux

func TestMain(m *testing.M) {
	app.InitSessionStore()
	testingMux = http.NewServeMux()
	NewAuthHandler(services.NewUserServiceMock()).SetRoutes(testingMux)

	m.Run()
}
