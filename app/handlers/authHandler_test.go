package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ZaphCode/F-SR-ChatApp/app/dtos"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
	"github.com/ZaphCode/F-SR-ChatApp/utils/mocks"
)

func TestSignUpCases(t *testing.T) {
	tests := []utils.AppHandlerTestCase[dtos.SignUpDto]{
		{
			Name: "Valid SignUp",
			Body: dtos.SignUpDto{
				Email: "omar@gmail.com", Password: "password", Username: "Omar",
			},
			ExpectedStatus: http.StatusCreated,
		},
		{
			Name:           "Invalid SignUp - Missing data",
			Body:           dtos.SignUpDto{},
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Name: "Invalid SignUp - Invalid email",
			Body: dtos.SignUpDto{
				Email: "invalid-email", Password: "password", Username: "Omar",
			},
			ExpectedStatus: http.StatusBadRequest,
		},
		{
			Name: "Invalid SignUp - Invalid password",
			Body: dtos.SignUpDto{
				Email: "omar@gmail.com", Password: "pass", Username: "Omar",
			},
			ExpectedStatus: http.StatusBadRequest,
		},
	}

	utils.RunAppHandlerTestCases(t, testingMux, http.MethodPost, "/api/auth/signup", tests)
}

func TestSuccessfulLoginFlow(t *testing.T) {
	body := fmt.Sprintf(`{"email": "%s","password": "%s"}`, mocks.UserA.Email, "password123")

	req, err := http.NewRequest("POST", "/api/auth/signin", strings.NewReader(body))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	testingMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	utils.PrettyPrint("Response: " + rr.Body.String())

	req, err = http.NewRequest("GET", "/api/auth/user", nil)

	if err != nil {
		t.Fatal(err)
	}

	for _, cookie := range rr.Result().Cookies() {
		req.AddCookie(cookie)
	}

	rr = httptest.NewRecorder()

	testingMux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	utils.PrettyPrint("Response: " + rr.Body.String())
}
