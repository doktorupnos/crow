package app

import (
	"context"
	"log"

	"github.com/doktorupnos/crow/backend/internal/env"
	"gorm.io/gorm"
)

// App groups all the state the server needs to run.
type App struct {
	Env *env.Env
	DB  *gorm.DB
}

func New(env *env.Env, db *gorm.DB) *App {
	return &App{
		Env: env,
		DB:  db,
	}
}

func (app *App) Run() {
	router := ConfiguredRouter(app.Env)
	server := GracefulServer(app, router)
	if err := server.ListenAndServe(context.Background()); err != nil {
		log.Fatal(err)
	}
	log.Println("graceful shutdown!")
}
