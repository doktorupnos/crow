package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/pressly/goose/v3"
)

func Run() {
	fmt.Println("Full Rewrite!")

	env, err := NewEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", env.DSN)
	if err != nil {
		log.Fatal(err)
	}

	goose.SetLogger(log.Default())
	if err := goose.Up(db, "sql/schema"); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}

	state := &State{
		db:        database.New(db),
		secret:    env.JWT.Secret,
		expiresIn: env.JWT.ExpiresIn,
	}
	router := Router(state)

	const addr = ":8000"
	server := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 1 * time.Second,
	}

	// TODO: graceful shutdown

	log.Println("api serving on", addr)
	err = server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
