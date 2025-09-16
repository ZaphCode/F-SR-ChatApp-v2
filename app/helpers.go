package app

import (
	"encoding/json"
	"net/http"

	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

//* JSON functions

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

// Rendering function (NOT NEEDED FOR NOW)
//// func Render(w http.ResponseWriter, tmplName string, data any) error {
//// 	tmplPath := filepath.Join("templates", tmplName+".html")
//// 	tmplFiles := []string{
//// 		tmplPath,
//// 		filepath.Join("templates", "components", "navbar.html"),
//// 	}
//// 	t, err := template.ParseFiles(tmplFiles...)
//// 	if err != nil {
//// 		return err
//// 	}
//// 	return t.ExecuteTemplate(w, tmplName, data)
//// }

//* HTTP Control types

type Handler interface {
	SetRoutes(mux *http.ServeMux)
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

//* Response types

type ApiStatus string

const StatusFail ApiStatus = "FAILURE"
const StatusSuccess ApiStatus = "SUCCESS"

type Response struct {
	Status ApiStatus `json:"status"`
	Msg    string    `json:"message"`
	Data   any       `json:"data,omitempty"`
	Error  any       `json:"error,omitempty"`
}

type JM map[string]any

//* Internal Context keys

type ContextKey string

const UserIDCtxKey ContextKey = "USER_ID"
