package integration

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/doktorupnos/crow/backend/internal/app"
	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/doktorupnos/crow/backend/internal/env"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const apiPrefix = "/api"

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
	ctx := context.Background()

	const dbUser = "postgres"
	const dbPassword = "postgres"
	const dbName = "crow"

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatalln("failed to start postgres container:", err)
	}
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatalln("failed to terminate postgres container:", err)
		}
	}()

	dsn, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalln("failed to get postgres container connection string:", err)
	}

	environment = &env.Env{
		Database: env.Database{
			DSN: dsn,
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

	db, err = database.Connect(environment.Database.DSN, &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		log.Println("Connect")
		log.Fatal(err)
		return
	}

	application = app.New(environment, db)
	router := chi.NewMux()
	router = app.RegisterEndpoints(router, application)
	server = httptest.NewServer(router)
	client = server.Client()

	usersEndpoint = server.URL + apiPrefix + "/users"
	loginEndpoint = server.URL + apiPrefix + "/login"
	logoutEndpoint = server.URL + apiPrefix + "/logout"
	validateJWTEndpoint = server.URL + apiPrefix + "/admin/jwt"
	postsEndpoint = server.URL + apiPrefix + "/posts"

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
