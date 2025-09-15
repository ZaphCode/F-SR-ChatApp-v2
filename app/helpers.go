package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/ZaphCode/F-SR-ChatApp/utils"
	"github.com/gorilla/sessions"
)

// Handler interface

type Handler interface {
	SetRoutes(mux *http.ServeMux)
}

// Json functions

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadAndValidateJson[T any](r *http.Request) (T, error) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, err
	}

	if err := utils.Validate(data); err != nil {
		return data, err
	}

	return data, nil
}

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

func Render(w http.ResponseWriter, tmplName string, data any) error {
	tmplPath := filepath.Join("templates", tmplName+".html")

	tmplFiles := []string{
		tmplPath,
		// More components...
		filepath.Join("templates", "components", "navbar.html"),
	}

	t, err := template.ParseFiles(tmplFiles...)

	if err != nil {
		return err
	}

	return t.ExecuteTemplate(w, tmplName, data)
}

type HandleFunc func(http.ResponseWriter, *http.Request) error

func (h HandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Middleware func(HandleFunc) HandleFunc

func (h HandleFunc) WithMiddlewares(middlewares ...Middleware) HandleFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// Json map helper

type JM map[string]any

// Responses

type ApiStatus string

const (
	StatusFail    ApiStatus = "FAILURE"
	StatusSuccess ApiStatus = "SUCCESS"
)

type Response struct {
	Status ApiStatus `json:"status"`
	Msg    string    `json:"message"`
	Data   any       `json:"data,omitempty"`
	Error  any       `json:"error,omitempty"`
}

type ContextKey string

const (
	UserIDCtxKey ContextKey = "USER_ID"
)
