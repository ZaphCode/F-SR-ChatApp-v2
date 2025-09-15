package handlers

import (
	"net/http"

	"github.com/ZaphCode/F-SR-ChatApp/domain"
)

type ConversationHandler struct {
	userService domain.UserService
}

func NewConversationHandler() *ConversationHandler {
	return &ConversationHandler{}
}

func (h *ConversationHandler) SetRoutes(mux *http.ServeMux) {

}

func (h *ConversationHandler) GetConversations(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting conversations
}
