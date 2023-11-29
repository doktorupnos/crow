package integration_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserWorkFlowIntegration(t *testing.T) {
  createUserServer := NewTestServer(application.CreateUser)
  defer createUserServer.Close()

  payload := strings.NewReader(`
  {
    "name": "user",
    "password": "password"
  }`)
  resp, err := http.Post(createUserServer.URL, "application/json", payload)
  if err != nil {
    t.Fatal(err)
  }
  defer resp.Body.Close()

  got := resp.StatusCode
  assertStatusCode(t, got, http.StatusCreated)

  // Try to create the same user again
  resp, err = http.Post(createUserServer.URL, "application/json", payload)
  if err != nil {
    t.Fatal(err)
  }
  defer resp.Body.Close()

  got = resp.StatusCode
  assertStatusCode(t, got, http.StatusBadRequest)


  loginServer := NewTestServer(application.BasicAuth(application.Login))
  defer loginServer.Close()
  
  loginRequest, err := http.NewRequest(http.MethodPost, loginServer.URL, strings.NewReader(``))
  if err != nil {
    t.Fatal(err)
  }
  loginRequest.SetBasicAuth("user", "password")

  loginResp, err := http.DefaultClient.Do(loginRequest)
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

  jwtValidateServer := NewTestServer(application.ValidateJWT)
  defer jwtValidateServer.Close()

  jwtValidateReq, err := http.NewRequest(http.MethodPost, jwtValidateServer.URL, strings.NewReader(``))
  if err != nil {
    t.Fatal(err)
  }
  jwtValidateReq.AddCookie(tokenCookie)

  jwtValidateResp, err := http.DefaultClient.Do(jwtValidateReq)
  if err != nil {
    return
  }
  defer jwtValidateResp.Body.Close()

  assertStatusCode(t, jwtValidateResp.StatusCode, http.StatusOK)

  updateUserServer := NewTestServer(application.JWT(application.UpdateUser))
  defer updateUserServer.Close()

  updateReq, err := http.NewRequest(http.MethodPut, updateUserServer.URL, strings.NewReader(`
    {
      "name": "updated_user",
      "password": "updated_password"
    }
  `))
  if err != nil {
    t.Fatal(err)
  }
  updateReq.AddCookie(tokenCookie)

  updateResp, err := http.DefaultClient.Do(updateReq)
  if err != nil {
    t.Fatal(err)
  }
  defer updateResp.Body.Close()

  assertStatusCode(t, updateResp.StatusCode, http.StatusOK)


  logoutServer := NewTestServer(application.JWT(application.Logout))
  defer logoutServer.Close()

  logoutReq, err := http.NewRequest(http.MethodPost, logoutServer.URL, strings.NewReader(``))
  if err != nil {
    t.Fatal(err)
  }
  logoutReq.AddCookie(tokenCookie)

  logoutResp, err := http.DefaultClient.Do(logoutReq)
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
  
  jwtValidateResp, err = http.DefaultClient.Do(jwtValidateReq)
  if err != nil {
    return
  }
  defer jwtValidateResp.Body.Close()

  assertStatusCode(t, jwtValidateResp.StatusCode, http.StatusOK)

  deleteUserServer := NewTestServer(application.BasicAuth(application.DeleteUser))
  defer deleteUserServer.Close()

  req, err := http.NewRequest(http.MethodDelete, deleteUserServer.URL, strings.NewReader(``))
  if err != nil {
    t.Fatal(err)
  }
  req.SetBasicAuth("updated_user", "updated_password")

  deleteResp, err := http.DefaultClient.Do(req)
  if err != nil {
    return 
  }
  defer deleteResp.Body.Close()

  assertStatusCode(t, deleteResp.StatusCode, http.StatusOK)
}

func TestUserCreateIntegeration(t *testing.T) {
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
