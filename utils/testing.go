package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestCase[T, R any] struct {
	Name         string
	Input        T
	ExpectError  bool
	HandleOutput func(t *testing.T, output R)
}

type AppHandlerTestCase[T any] struct {
	Name           string
	Body           T
	ExpectedStatus int
}

func RunAppHandlerTestCases[T any](
	t *testing.T,
	mux *http.ServeMux,
	method, endpoint string,
	testCases []AppHandlerTestCase[T],
) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			b, _ := json.Marshal(tc.Body)

			req, _ := http.NewRequest(method, endpoint, strings.NewReader(string(b)))

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.ExpectedStatus)
			}

			PrettyPrint("Response " + getResponseJson(rr.Body.Bytes()))
		})
	}
}

func getResponseJson(input []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, input, "", "  ")
	if err != nil {
		return string(input)
	}
	return out.String()
}
