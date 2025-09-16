package handlers

import (
	"net/http"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/ZaphCode/F-SR-ChatApp/app/dtos"
	"github.com/ZaphCode/F-SR-ChatApp/app/middlewares"
	"github.com/ZaphCode/F-SR-ChatApp/domain"
	"github.com/google/uuid"
)

type AuthHandler struct {
	us domain.UserService
}

func NewAuthHandler(userService domain.UserService) *AuthHandler {
	return &AuthHandler{
		us: userService,
	}
}

// Routes

func (h *AuthHandler) SetRoutes(mux *http.ServeMux) {
	mux.Handle("POST /api/auth/signup", app.HandleFunc(h.SignUp))
	mux.Handle("POST /api/auth/signin", app.HandleFunc(h.SignIn))
	mux.Handle("POST /api/auth/signout", app.HandleFunc(h.SignOut))
	mux.Handle("GET /api/auth/user", app.HandleFunc(h.GetAuthUser).WithMiddlewares(middlewares.Auth))
}

// Handlers

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	newUser, err := app.ReadAndValidateJson[dtos.SignUpDto](r)

	if err != nil {
		return app.WriteJson(w, http.StatusBadRequest, app.Response{
			Status: app.StatusFail, Msg: "Invalid data", Error: err,
		})
	}

	user, err := h.us.Create(newUser.Username, newUser.Email, newUser.Password)

	if err != nil {
		return app.WriteJson(w, http.StatusBadRequest, app.Response{
			Status: app.StatusFail, Msg: "Error while creating user", Error: err.Error(),
		})
	}

	if err := app.SaveSessionValue(w, r, "user_id", user.ID.String()); err != nil {
		return app.WriteJson(w, http.StatusInternalServerError, app.Response{
			Status: app.StatusFail, Msg: "Error while saving session", Error: err.Error(),
		})
	}

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

	user, err := h.us.Authenticate(credentials.Email, credentials.Password)

	if err != nil {
		return app.WriteJson(w, http.StatusBadRequest, app.Response{
			Status: app.StatusFail, Msg: "Invalid credentials", Error: err,
		})
	}

	if err := app.SaveSessionValue(w, r, "user_id", user.ID.String()); err != nil {
		return app.WriteJson(w, http.StatusInternalServerError, app.Response{
			Status: app.StatusFail, Msg: "Error while saving session", Error: err.Error(),
		})
	}

	return app.WriteJson(w, http.StatusOK, app.Response{
		Status: app.StatusSuccess, Msg: "User signed in successfully", Data: user,
	})
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) error {
	if err := app.DeleteSessionValue(w, r, "user_id"); err != nil {
		return app.WriteJson(w, http.StatusInternalServerError, app.Response{
			Status: app.StatusFail, Msg: "Error while signing out", Error: err.Error(),
		})
	}

	return app.WriteJson(w, http.StatusOK, app.Response{
		Status: app.StatusSuccess, Msg: "User signed out successfully",
	})
}

func (h *AuthHandler) GetAuthUser(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(app.UserIDCtxKey).(uuid.UUID)

	user, err := h.us.GetByID(userID)

	if err != nil {
		return app.WriteJson(w, http.StatusInternalServerError, app.Response{
			Status: app.StatusFail, Msg: "Something went wrong", Error: err,
		})
	}

	return app.WriteJson(w, http.StatusOK, app.Response{
		Status: app.StatusSuccess, Msg: "User retrieved successfully", Data: user,
	})
}
