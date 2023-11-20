package main

import (
	"log"
	"net/http"
	"os"

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
	// .env will only be used for local development
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env : %q", err)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT environment variable is not set")
	}

	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		log.Fatal("DSN environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database\nData Source Name : %q\nError : %q", dsn, err)
	}

	ping(db)
	migrate(db)

	cfg := &ApiConfig{DB: db}

	mainRouter := chi.NewRouter()
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
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

	userRouter := chi.NewRouter()
	userRouter.Post("/", cfg.CreateUser)
	userRouter.Get("/", cfg.GetAllUsers)

	mainRouter.Mount("/users", userRouter)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	log.Print("Serving on port:", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func ping(db *gorm.DB) {
	sqlDB, err := db.DB()
	// Realistically this shouldn't fail
	if err != nil {
		log.Fatalf("Failed to get generic *sql.DB : %q", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Database ping failed : %q", err)
	}
}

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Failed to perform migration on the User type : %q", err)
	}
}
