package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserWorkFlowIntegration(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	token := createTestUser(t, "user", "password", http.StatusCreated)
	validateJwt(t, token)

	createTestUser(t, "user", "password", http.StatusBadRequest)

	login(t, "user", "password")

	updateReq, err := http.NewRequest(http.MethodPut, usersEndpoint, strings.NewReader(`
    {
      "name": "updated_user",
      "password": "updated_password"
    }
  `))
	if err != nil {
		t.Fatal(err)
	}
	updateReq.AddCookie(token)

	updateResp, err := client.Do(updateReq)
	if err != nil {
		t.Fatal(err)
	}
	defer updateResp.Body.Close()
	assertStatusCode(t, updateResp.StatusCode, http.StatusOK)

	logout(t, token)

	deleteTestUser(t, "updated_user", "updated_password")
}

func TestUserCreateIntegeration(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	server := httptest.NewServer(http.HandlerFunc(application.CreateUser))
	defer server.Close()

	t.Run("empty body", func(t *testing.T) {
		createTestUser(t, "", "", http.StatusBadRequest)
	})

	t.Run("empty name", func(t *testing.T) {
		createTestUser(t, "", "password", http.StatusBadRequest)
	})

	t.Run("empty password", func(t *testing.T) {
		createTestUser(t, "name", "", http.StatusBadRequest)
	})

	t.Run("name too big", func(t *testing.T) {
		createTestUser(t, "a_valid_name_over_twenty_characters", "password", http.StatusBadRequest)
	})

	t.Run("password too big", func(t *testing.T) {
		password := strings.Repeat("a", 80)
		createTestUser(t, "name", password, http.StatusBadRequest)
	})

	t.Run("name contains not supported characters", func(t *testing.T) {
		name := "not- All*w$d"
		createTestUser(t, name, "password", http.StatusBadRequest)
	})
}

func createTestUser(t testing.TB, name, password string, expectedStatusCode int) *http.Cookie {
	t.Helper()

	payload := strings.NewReader(fmt.Sprintf(`{
    "name": %q,
    "password": %q
  }`, name, password))
	createResp, err := client.Post(usersEndpoint, "application/json", payload)
	if err != nil {
		t.Fatal(err)
	}
	defer createResp.Body.Close()

	got := createResp.StatusCode
	assertStatusCode(t, got, expectedStatusCode)

	if expectedStatusCode != http.StatusCreated {
		return nil
	}
	return extractTokenCookie(t, createResp)
}

func extractTokenCookie(t testing.TB, resp *http.Response) *http.Cookie {
	t.Helper()

	for _, c := range resp.Cookies() {
		if c.Name == "token" {
			return c
		}
	}

	t.Fatal("token cookie was not set")
	return nil
}

func deleteTestUser(t testing.TB, name, password string) {
	t.Helper()

	deleteReq, err := http.NewRequest(http.MethodDelete, usersEndpoint, noBody)
	if err != nil {
		t.Fatal(err)
	}
	deleteReq.SetBasicAuth(name, password)

	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		return
	}
	defer deleteResp.Body.Close()

	assertStatusCode(t, deleteResp.StatusCode, http.StatusOK)
}

func validateJwt(t testing.TB, token *http.Cookie) {
	t.Helper()

	jwtValidateReq, err := http.NewRequest(http.MethodPost, validateJWTEndpoint, noBody)
	if err != nil {
		t.Fatal(err)
	}
	jwtValidateReq.AddCookie(token)

	jwtValidateResp, err := client.Do(jwtValidateReq)
	if err != nil {
		return
	}
	defer jwtValidateResp.Body.Close()

	assertStatusCode(t, jwtValidateResp.StatusCode, http.StatusOK)
}

func login(t testing.TB, username, passwrod string) *http.Cookie {
	loginRequest, err := http.NewRequest(http.MethodPost, loginEndpoint, noBody)
	if err != nil {
		t.Fatal(err)
	}
	loginRequest.SetBasicAuth("user", "password")

	loginResp, err := client.Do(loginRequest)
	if err != nil {
		t.Fatal(err)
	}
	defer loginResp.Body.Close()
	assertStatusCode(t, loginResp.StatusCode, http.StatusOK)

	token := extractTokenCookie(t, loginResp)
	validateJwt(t, token)
	return token
}

func logout(t testing.TB, token *http.Cookie) {
	logoutReq, err := http.NewRequest(http.MethodPost, logoutEndpoint, noBody)
	if err != nil {
		t.Fatal(err)
	}
	logoutReq.AddCookie(token)

	logoutResp, err := client.Do(logoutReq)
	if err != nil {
		t.Fatal(err)
	}
	defer logoutResp.Body.Close()
	assertStatusCode(t, logoutResp.StatusCode, http.StatusOK)
}
