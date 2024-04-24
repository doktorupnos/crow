package app

import (
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/doktorupnos/crow/backend/internal/env"
	"github.com/doktorupnos/crow/backend/internal/follow"
	"github.com/doktorupnos/crow/backend/internal/like"
	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/doktorupnos/crow/backend/internal/shutdown"
	"github.com/doktorupnos/crow/backend/internal/user"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type App struct {
	Env           *env.Env
	DB            *gorm.DB
	userService   *user.Service
	postService   *post.Service
	likeService   *like.Service
	followService *follow.Service
}

func New(env *env.Env, db *gorm.DB) *App {
	us := user.NewService(database.NewGormUserRepo(db))
	ps := post.NewService(database.NewGormPostRepo(db))
	ls := like.NewService(database.NewGormLikeRepo(db))
	fs := follow.NewService(database.NewGormFollowRepo(db), us)
	return &App{
		Env:           env,
		DB:            db,
		userService:   us,
		postService:   ps,
		likeService:   ls,
		followService: fs,
	}
}

func (app *App) Run() {
	router := chi.NewMux()
	RegisterMiddleware(router, app)
	router = RegisterEndpoints(router, app)
	server := &http.Server{
		Addr:    app.Env.Server.Addr,
		Handler: router,

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}
	shutdown.ListenAndServe(server, app.Env.Server.ShutdownTimeout)
}
