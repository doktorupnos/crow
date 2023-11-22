package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ApiConfig struct {
	DB *gorm.DB
}

func main() {
	local := flag.Bool("local", false, "Depend on a .env file for local development")
	flag.Parse()

	if *local {
		err := godotenv.Load()
		if err != nil {
			log.Printf("ERROR: failed to load .env : %q", err)
		}
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("ERROR: PORT environment variable is not set")
	}

	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		log.Fatal("ERROR: DSN environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(
			"ERROR: Failed to connect to database\nData Source Name : %q\nError : %q",
			dsn,
			err,
		)
	}

	ping(db)
	migrate(db)

	cfg := &ApiConfig{DB: db}

	mainRouter := chi.NewRouter()
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))
	mainRouter.Use(middleware.Logger)

	mainRouter.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		type StatusResponse struct {
			Status string `json:"status"`
		}
		// TODO: enhance to return a http.StatusServiceUnavailable.
		respondWithJSON(
			w,
			http.StatusOK,
			StatusResponse{Status: http.StatusText(http.StatusOK)},
		)
	})

	userRouter := chi.NewRouter()
	userRouter.Route("/", func(r chi.Router) {
		r.Post("/", cfg.CreateUser)
		r.Get("/", cfg.GetAllUsers)
	})
	mainRouter.Mount("/users", userRouter)

	// TODO: gracefully shutdown the server for k8s.
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,

		// Good practice to set timeouts to avoid Slowloris attacks.
		// The Timeout values are not final.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("STATUS: serving on :%s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func ping(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf(
			"failed to get generic *sql.DB while trying to ping the database\nError : %q",
			err,
		)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("ERROR: database ping failed : %q", err)
	}
}

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to perform migration on the User type : %q", err)
	}
}
