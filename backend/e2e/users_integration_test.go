package integration_test

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

  payload := strings.NewReader(`
  {
    "name": "user",
    "password": "password"
  }`)
  resp, err := client.Post(usersEndpoint, "application/json", payload)
  if err != nil {
    t.Fatal(err)
  }
  defer resp.Body.Close()

  got := resp.StatusCode
  assertStatusCode(t, got, http.StatusCreated)

  // Try to create the same user again
  resp, err = client.Post(usersEndpoint, "application/json", payload)
  if err != nil {
    t.Fatal(err)
  }
  defer resp.Body.Close()

  got = resp.StatusCode
  assertStatusCode(t, got, http.StatusBadRequest)

  loginRequest, err := http.NewRequest(http.MethodPost, loginEndpoint, strings.NewReader(``))
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

  cookies := loginResp.Cookies()
  var tokenCookie *http.Cookie
  for _, c := range cookies {
    if c.Name == "token" {
      tokenCookie = c
      break
    }
  }
  if tokenCookie == nil {
    t.Fatal("token cookie was not set")
  }
  t.Logf("%s : %s", tokenCookie.Name, tokenCookie.Value)

  jwtValidateReq, err := http.NewRequest(http.MethodPost, validateJWTEndpoint, strings.NewReader(``))
  if err != nil {
    t.Fatal(err)
  }
  jwtValidateReq.AddCookie(tokenCookie)

  jwtValidateResp, err := client.Do(jwtValidateReq)
  if err != nil {
    return
  }
  defer jwtValidateResp.Body.Close()

  assertStatusCode(t, jwtValidateResp.StatusCode, http.StatusOK)

  updateReq, err := http.NewRequest(http.MethodPut, usersEndpoint, strings.NewReader(`
    {
      "name": "updated_user",
      "password": "updated_password"
    }
  `))
  if err != nil {
    t.Fatal(err)
  }
  updateReq.AddCookie(tokenCookie)

  updateResp, err := client.Do(updateReq)
  if err != nil {
    t.Fatal(err)
  }
  defer updateResp.Body.Close()

  assertStatusCode(t, updateResp.StatusCode, http.StatusOK)


  logoutReq, err := http.NewRequest(http.MethodPost, logoutEndpoint, strings.NewReader(``))
  if err != nil {
    t.Fatal(err)
  }
  logoutReq.AddCookie(tokenCookie)

  logoutResp, err := client.Do(logoutReq)
  if err != nil {
    t.Fatal(err)
  }
  defer logoutResp.Body.Close()

  assertStatusCode(t, logoutResp.StatusCode, http.StatusOK)

  for _, c := range logoutResp.Cookies(){
    if c.Name == "token" {
      tokenCookie = c
      break
    }
  }
  if tokenCookie == nil {
    t.Fatal("token cookie was not set")
  }
  jwtValidateReq.AddCookie(tokenCookie)
  
  jwtValidateResp, err = client.Do(jwtValidateReq)
  if err != nil {
    return
  }
  defer jwtValidateResp.Body.Close()

  assertStatusCode(t, jwtValidateResp.StatusCode, http.StatusOK)

  deleteReq, err := http.NewRequest(http.MethodDelete, usersEndpoint, strings.NewReader(``))
  if err != nil {
    t.Fatal(err)
  }
  deleteReq.SetBasicAuth("updated_user", "updated_password")

  deleteResp, err := client.Do(deleteReq)
  if err != nil {
    return 
  }
  defer deleteResp.Body.Close()

  assertStatusCode(t, deleteResp.StatusCode, http.StatusOK)
}

func TestUserCreateIntegeration(t *testing.T) {
  if testing.Short() {
    t.SkipNow()
  }

  server := httptest.NewServer(http.HandlerFunc(application.CreateUser))
  defer server.Close()

  t.Run("empty body", func(t *testing.T) {

    payload := strings.NewReader(``)
    resp, err := http.Post(server.URL, "application/json", payload)
    if err != nil {
      t.Fatal(err)
    }
    defer resp.Body.Close()

    got := resp.StatusCode
    assertStatusCode(t, got, http.StatusBadRequest)
  })

  t.Run("empty name", func(t *testing.T) {
    payload := strings.NewReader(`
    {
      "password": "password"
    }`)
    resp, err := http.Post(server.URL, "application/json", payload)
    if err != nil {
      t.Fatal(err)
    }
    defer resp.Body.Close()

    got := resp.StatusCode
    assertStatusCode(t, got, http.StatusBadRequest)
  })

  t.Run("empty password", func(t *testing.T) {
    payload := strings.NewReader(`
    {
      "name": "name"
    }`)
    resp, err := http.Post(server.URL, "application/json", payload)
    if err != nil {
      t.Fatal(err)
    }
    defer resp.Body.Close()

    got := resp.StatusCode
    assertStatusCode(t, got, http.StatusBadRequest)
  })

  t.Run("name too big", func(t *testing.T) {
    payload := strings.NewReader(`
    {
      "name": "a_valid_name_over_twenty_characters",
      "password": "password"
    }`)
    resp, err := http.Post(server.URL, "application/json", payload)
    if err != nil {
      t.Fatal(err)
    }
    defer resp.Body.Close()

    got := resp.StatusCode
    assertStatusCode(t, got, http.StatusBadRequest)
  })

  t.Run("password too big", func(t *testing.T) {
    password := strings.Repeat("a", 80)
    payload := strings.NewReader(fmt.Sprintf(`
    {
      "name": "name",
      "password": "%s"
    }`, password))
    resp, err := http.Post(server.URL, "application/json", payload)
    if err != nil {
      t.Fatal(err)
    }
    defer resp.Body.Close()

    got := resp.StatusCode
    assertStatusCode(t, got, http.StatusBadRequest)
  })

  t.Run("name contains not supported characters", func(t *testing.T) {
    payload := strings.NewReader(`
    {
      "name": "not- All*w$d",
      "password": "password"
    }`)
    resp, err := http.Post(server.URL, "application/json", payload)
    if err != nil {
      t.Fatal(err)
    }
    defer resp.Body.Close()

    got := resp.StatusCode
    assertStatusCode(t, got, http.StatusBadRequest)
  })
}
