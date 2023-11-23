package main

import (
	"log"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/database"
)

func main() {
	env, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(env.DSN)
	if err != nil {
		log.Fatal(err)
	}

	app := &App{
		DB:         db,
		JWT_SECRET: env.DSN,
	}

	router := registerRoutes(app)

	// TODO: gracefully shutdown the server for k8s.
	// NOTE: When gracefull shutdown is implemented the server definition should have its own function.
	server := &http.Server{
		Addr:    ":" + env.PORT,
		Handler: router,

		// Good practice to set timeouts to avoid Slowloris attacks.
		// The Timeout values are not final and should be tested.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("STATUS: serving on :%s", env.PORT)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
