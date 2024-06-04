package app

import (
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/channel"
	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/doktorupnos/crow/backend/internal/env"
	"github.com/doktorupnos/crow/backend/internal/follow"
	"github.com/doktorupnos/crow/backend/internal/like"
	"github.com/doktorupnos/crow/backend/internal/message"
	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/doktorupnos/crow/backend/internal/shutdown"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/gorilla/websocket"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type App struct {
	Env            *env.Env
	DB             *gorm.DB
	chatServer     *ChatServer
	userService    *user.Service
	postService    *post.Service
	likeService    *like.Service
	followService  *follow.Service
	channelService *channel.Service
	upgrader       websocket.Upgrader
}

func New(env *env.Env, db *gorm.DB) *App {
	us := user.NewService(database.NewGormUserRepo(db))
	return &App{
		Env: env,
		DB:  db,
		chatServer: NewChatServer(
			message.NewService(database.NewGormMessageRepo(db)),
		),
		userService:    us,
		postService:    post.NewService(database.NewGormPostRepo(db)),
		likeService:    like.NewService(database.NewGormLikeRepo(db)),
		followService:  follow.NewService(database.NewGormFollowRepo(db), us),
		channelService: channel.NewService(database.NewGormChannelRepo(db)),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
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
