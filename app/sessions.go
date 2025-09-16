package app

import (
	"fmt"
	"net/http"

	"github.com/ZaphCode/F-SR-ChatApp/utils"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func InitSessionStore() {
	store = sessions.NewCookieStore([]byte(utils.APP_SESSION_KEY))
	store.Options.HttpOnly = true
	store.Options.Secure = false
	store.Options.SameSite = http.SameSiteDefaultMode
	store.Options.MaxAge = 3600 * 24 * 7 // 1 week
}

func SaveSessionValue(w http.ResponseWriter, r *http.Request, key string, value any) error {
	session, err := store.Get(r, utils.APP_SESSION_COOKIE)

	if err != nil {
		return err
	}

	session.Values[key] = value

	return session.Save(r, w)
}

func GetSessionValue[T any](r *http.Request, key string) (T, error) {
	session, err := store.Get(r, utils.APP_SESSION_COOKIE)

	if err != nil {
		return *new(T), err
	}

	val, exists := session.Values[key]

	if !exists {
		return *new(T), nil
	}

	typedVal, ok := val.(T)

	if !ok {
		return *new(T), fmt.Errorf("invalid type assertion for key %s", key)
	}

	return typedVal, nil
}

func DeleteSessionValue(w http.ResponseWriter, r *http.Request, key string) error {
	session, err := store.Get(r, utils.APP_SESSION_COOKIE)

	if err != nil {
		return err
	}

	delete(session.Values, key)

	return session.Save(r, w)
}
