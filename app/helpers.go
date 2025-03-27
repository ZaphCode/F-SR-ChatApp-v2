package app

import (
	"encoding/json"
	"net/http"
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

	if err := validate(data); err != nil {
		return data, err
	}

	return data, nil
}

// Custom handler function

type HandleFunc func(http.ResponseWriter, *http.Request) error

func (h HandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	Status  ApiStatus `json:"status"`
	Message string    `json:"message"`
	Data    any       `json:"data,omitempty"`
	Error   any       `json:"error,omitempty"`
}
