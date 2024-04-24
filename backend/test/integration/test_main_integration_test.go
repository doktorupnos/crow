package integration

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/doktorupnos/crow/backend/internal/app"
	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/doktorupnos/crow/backend/internal/env"
	"gorm.io/gorm"
)

var (
	environment *env.Env
	db          *gorm.DB
	application *app.App
	router      http.Handler
	server      *httptest.Server
	client      *http.Client
)

var (
	usersEndpoint       string
	loginEndpoint       string
	logoutEndpoint      string
	validateJWTEndpoint string
	postsEndpoint       string
)

var noBody = strings.NewReader(``)

func TestMain(m *testing.M) {
	environment = &env.Env{
		Database: env.Database{
			DSN: `postgres://postgres:postgres@localhost:5432/crow`,
		},
		JWT: env.JWT{
			Secret:   "+3xObWCCIAQf/N1ltJD27kZ5gfjmfbUBG4ViZ/6oHI3rpVFmhAo7yzwWg4mivB1Jea8UuwooegxTdZhZgLkZZA==",
			Lifetime: 5 * time.Minute,
		},
		Pagination: env.Pagination{
			DefaultPostsPageSize:  3,
			DefaultFollowPageSize: 5,
		},
		Posts: env.Posts{
			BodyLimit: 280,
		},
	}

	var err error
	db, err = database.Connect(environment.Database.DSN)
	if err != nil {
		log.Fatal(err)
		return
	}

	application = app.New(environment, db)
	router = app.ConfiguredRouter(application)
	server = httptest.NewServer(router)
	client = server.Client()

	usersEndpoint = server.URL + "/users"
	loginEndpoint = server.URL + "/login"
	logoutEndpoint = server.URL + "/logout"
	validateJWTEndpoint = server.URL + "/admin/jwt"
	postsEndpoint = server.URL + "/posts"

	defer server.Close()
	m.Run()
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
		t.Fatalf("\ngot: %s\nwant: %s", http.StatusText(got), http.StatusText(want))
	}
}
