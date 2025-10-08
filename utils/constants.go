package utils

import "errors"

// App session
const APP_SESSION_KEY = "f-sr-session-key"
const APP_SESSION_COOKIE = "fsr-session"

// Database
const (
	MONGO_USER_COL         = "users"
	MONGO_CONVERSATION_COL = "conversations"
	MONGO_MESSAGE_COL      = "messages"
	MONGO_DEV_URI          = "mongodb://zaph:zaphpass@localhost:27017/"
	MONGO_DEV_DB           = "fsr-chatapp-dev"
)

// Errors
var ErrNotFound = errors.New("not found")
