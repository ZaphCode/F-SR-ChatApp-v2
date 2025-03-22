package handlers

import (
	"net/http"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/ZaphCode/F-SR-ChatApp/app/dtos"
	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService domain.UserService
}

// TODO: Add service params
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h UserHandler) SetRoutes(mux *http.ServeMux) {
	mux.Handle("GET /api/users", app.HandleFunc(h.GetUsers))
	mux.Handle("GET /api/user", app.HandleFunc(h.GetUser))
	mux.Handle("POST /api/user", app.HandleFunc(h.CreateUser))
}

// Handlers

func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	return app.WriteJson(w, http.StatusOK, app.JM{"users": []domain.User{
		{ID: uuid.New(), Username: "Alice"},
		{ID: uuid.New(), Username: "Bob"},
	}})
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) error {
	return app.WriteJson(w, http.StatusOK, map[string]any{
		"user": domain.User{ID: uuid.New(), Username: "Alice"},
	})
}

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	body, err := app.ReadAndValidateJson[dtos.NewUserDto](r)

	if err != nil {
		return app.WriteJson(w, http.StatusBadRequest, app.JM{"error": err})
	}

	user := body.AdaptToUser()

	if err := h.userService.Create(&user); err != nil {
		return app.WriteJson(w, http.StatusBadRequest, app.JM{"error": err})
	}

	return app.WriteJson(w, http.StatusCreated, app.JM{"user": user})
}
