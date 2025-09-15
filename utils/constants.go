package utils

import "errors"

// App session
const APP_SESSION_KEY = "f-sr-session-key"
const APP_SESSION_COOKIE = "fsr-session"

// Development database uri
const MONGO_DEV_URI = "mongodb://zaph:zaphpass@localhost:27017/"

// Errors
var ErrNotFound = errors.New("not found")
