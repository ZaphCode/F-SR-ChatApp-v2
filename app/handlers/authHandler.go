package handlers

import (
	"net/http"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/ZaphCode/F-SR-ChatApp/app/dtos"
	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

type AuthHandler struct {
	userService domain.UserService
}

func NewAuthHandler(userService domain.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// Routes

func (h *AuthHandler) SetRoutes(mux *http.ServeMux) {
	mux.Handle("GET /signup", app.HandleFunc(h.SignUpView))
}

// Handlers

func (h *AuthHandler) SignUpView(w http.ResponseWriter, r *http.Request) error {
	return app.Render(w, "sign-up", nil)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	newUser, err := app.ReadAndValidateJson[dtos.SignUpDto](r)

	if err != nil {
		return app.WriteJson(w, http.StatusBadRequest, app.Response{
			Status: app.StatusFail, Msg: "Invalid data", Error: err,
		})
	}

	user, err := h.userService.Create(newUser.Username, newUser.Email, newUser.Password)

	if err != nil {
		return app.WriteJson(w, http.StatusBadRequest, app.Response{
			Status: app.StatusFail, Msg: "Something went wrong while creating user", Error: err,
		})
	}

	app.SaveSessionValue(w, r, "user_id", user.ID)

	return app.WriteJson(w, http.StatusCreated, app.Response{
		Status: app.StatusSuccess, Msg: "User created successfully", Data: user,
	})
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) error {
	credentials, err := app.ReadAndValidateJson[dtos.SignInDto](r)

	if err != nil {
		return app.WriteJson(w, http.StatusBadRequest, app.Response{
			Status: app.StatusFail, Msg: "Invalid data", Error: err,
		})
	}

	user, err := h.userService.Authenticate(credentials.Email, credentials.Password)

	if err != nil {
		return app.WriteJson(w, http.StatusBadRequest, app.Response{
			Status: app.StatusFail, Msg: "Invalid credentials", Error: err,
		})
	}

	app.SaveSessionValue(w, r, "user_id", user.ID)

	return app.WriteJson(w, http.StatusCreated, app.Response{
		Status: app.StatusSuccess, Msg: "User signed in successfully", Data: user,
	})
}
