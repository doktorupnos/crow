package app

import (
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/doktorupnos/crow/backend/internal/env"
	"github.com/doktorupnos/crow/backend/internal/shutdown"
	"gorm.io/gorm"
)

// App groups all the state the server needs to run.
type App struct {
	Env             *env.Env
	DB              *gorm.DB
	userService     *UserService
	postService     *PostService
	postLikeService *PostLikeService
}

func New(env *env.Env, db *gorm.DB) *App {
	return &App{
		Env:             env,
		DB:              db,
		userService:     NewUserService(database.NewGormUserRepo(db)),
		postService:     NewPostService(database.NewGormPostRepo(db)),
		postLikeService: NewPostLikeService(database.NewGormPostLikeRepo(db)),
	}
}

func (app *App) Run() {
	router := ConfiguredRouter(app)
	server := &http.Server{
		Addr:    app.Env.Server.Addr,
		Handler: router,

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}
	shutdown.ListenAndServe(server, app.Env.Server.ShutdownTimeout)
}
