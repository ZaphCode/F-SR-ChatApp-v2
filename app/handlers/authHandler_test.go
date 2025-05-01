package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/ZaphCode/F-SR-ChatApp/app/dtos"
	"github.com/ZaphCode/F-SR-ChatApp/services"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

var mux *http.ServeMux

func TestMain(m *testing.M) {
	app.InitSessionStore()
	mux = http.NewServeMux()
	NewAuthHandler(services.NewUserServiceMock()).SetRoutes(mux)

	m.Run()
}

func TestSignUpCases(t *testing.T) {
	tests := []utils.TestAppHandlerCase[dtos.SignUpDto]{
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

	utils.RunTestCases(t, mux, http.MethodPost, "/api/auth/signup", tests)
}

func TestSuccessfulLoginFlow(t *testing.T) {
	body := `{"email":"test@user.com","password":"testpassword"}`

	req, err := http.NewRequest("POST", "/api/auth/signin", strings.NewReader(body))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	t.Logf("Response: %s", rr.Body.String())

	req, err = http.NewRequest("GET", "/api/auth/user", nil)

	if err != nil {
		t.Fatal(err)
	}

	for _, cookie := range rr.Result().Cookies() {
		req.AddCookie(cookie)
	}

	rr = httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	t.Logf("\033[34mResponse: %s\033[0m", rr.Body.String())

}
