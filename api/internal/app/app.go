package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/pressly/goose/v3"
)

func Run() {
	fmt.Println("Full Rewrite!")

	// TODO: environment
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		log.Fatal("DSN environment variable is not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	goose.SetLogger(log.Default())
	if err := goose.Up(db, "sql/schema"); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}

	state := &State{
		db:        database.New(db),
		secret:    "makaronia",
		expiresIn: time.Hour,
	}
	router := Router(state)

	const addr = ":8000"
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// TODO: graceful shutdown

	log.Println("api serving on", addr)
	err = server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
