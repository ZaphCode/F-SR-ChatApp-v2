package handlers

import (
	"net/http"

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
	mux.Handle("POST /api/conversations", app.HandleFunc(h.CreateConversation))
}

// Handlers

func (h *ConversationHandler) GetConversations(w http.ResponseWriter, r *http.Request) error {
	// Implementation for getting conversations
	return nil
}

func (h *ConversationHandler) CreateConversation(w http.ResponseWriter, r *http.Request) error {
	return nil
}
