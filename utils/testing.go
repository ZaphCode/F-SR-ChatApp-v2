package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestAppHandlerCase[T any] struct {
	Name           string
	Body           T
	ExpectedStatus int
}

func RunTestCases[T any](t *testing.T, mux *http.ServeMux, method, url string, testCases []TestAppHandlerCase[T]) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			b, _ := json.Marshal(tc.Body)

			req, _ := http.NewRequest(method, url, strings.NewReader(string(b)))

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.ExpectedStatus)
			}

			// print body blue colored
			t.Logf("\033[34m\n--- Response: %s\033[0m", rr.Body.String())
		})
	}
}
