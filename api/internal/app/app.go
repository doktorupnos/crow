package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/api/internal/database"
)

func Run() {
	fmt.Println("Full Rewrite!")

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/crow?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	state := &State{
		DB:        database.New(db),
		Secret:    "makaronia",
		ExpiresIn: time.Hour,
	}
	router := Router(state)

	const addr = ":8000"
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("api serving on", addr)
	err = server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
