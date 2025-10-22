package handlers

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

type ConversationHandler struct {
	us domain.UserService
	cs domain.ConversationService
}

func NewConversationHandler(
	userService domain.UserService,
	conversationService domain.ConversationService,
) *ConversationHandler {
	return &ConversationHandler{
		us: userService,
		cs: conversationService,
	}
}

// Routes

func (h *ConversationHandler) SetRoutes(mux *http.ServeMux) {
	mux.Handle("GET /api/conversations", app.HandleFunc(h.GetConversations))
	mux.Handle("POST /api/conversation", app.HandleFunc(h.GetOrCreateConversation))
}

// Handlers

func (h *ConversationHandler) GetConversations(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(app.UserIDCtxKey).(uuid.UUID)

	conversations, err := h.cs.GetAllFrom(userID)

	if err != nil {
		return app.WriteJson(w, http.StatusInternalServerError, app.Response{
			Status: app.StatusFail, Msg: "Something went wrong", Error: err,
		})
	}

	return app.WriteJson(w, http.StatusOK, app.Response{
		Status: app.StatusSuccess, Msg: "Conversations retrieved successfully", Data: conversations,
	})
}

func (h *ConversationHandler) GetOrCreateConversation(w http.ResponseWriter, r *http.Request) error {
	return nil
}
