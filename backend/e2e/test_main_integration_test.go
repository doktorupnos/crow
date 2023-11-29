package integration_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/doktorupnos/crow/backend/internal/app"
	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/doktorupnos/crow/backend/internal/env"
)

var application *app.App

func TestMain(m *testing.M) {
	env := &env.Env{}
	env.DSN = `postgres://postgres:postgres@localhost:5432/crow`
	env.JwtSecret = "+3xObWCCIAQf/N1ltJD27kZ5gfjmfbUBG4ViZ/6oHI3rpVFmhAo7yzwWg4mivB1Jea8UuwooegxTdZhZgLkZZA=="
	env.JwtLifetime = 5 * time.Minute

	db, err := database.Connect(env.DSN)
	if err != nil {
		log.Fatal(err)
	}

	application = app.New(env, db)
	m.Run()
}

func NewTestServer(h http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(h))
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Fatalf("\ngot: %+v\nwant: %+v", got, want)
	}
}

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot: %s\nwant: %s", http.StatusText(got), http.StatusText(want))
	}
}
