package app

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func RegisterEndpoints(r *chi.Mux, app *App) *chi.Mux {
	apiRouter := chi.NewRouter()

	r.Get("/healthz", HealthCheck)

	r.Post("/login", app.BasicAuth(app.Login))
	r.Post("/logout", app.JWT(app.Logout))

	r.Get("/profile", app.JWT(app.ViewProfile))

	r.Post("/follow", app.JWT(app.Follow))
	r.Post("/unfollow", app.JWT(app.UnFollow))
	r.Get("/following", app.JWT(app.Following))
	r.Get("/followers", app.JWT(app.Followers))
	r.Get("/following_count", app.JWT(app.FollowingCount))
	r.Get("/followers_count", app.JWT(app.FollowerCount))

	r.Mount("/users", UserRouter(app))
	r.Mount("/posts", PostRouter(app))
	r.Mount("/post_likes", PostLikeRouter(app))
	r.Mount("/admin", AdminRouter(app))

	r.Handle("/chat", app.WS(app.chatServer.cosmos))

	r.Handle("/ws/echo", app.chatServer.upgrade(app.chatServer.echo))
	r.Handle("/ws/feed", app.chatServer.upgrade(app.chatServer.feed))

	apiRouter.Mount("/api", r)
	return apiRouter
}

func RegisterMiddleware(router *chi.Mux, app *App) {
	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins: []string{app.Env.Server.CorsOrigin},
			AllowedHeaders: []string{"*"},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodOptions,
				http.MethodPut,
				http.MethodDelete,
			},
			AllowCredentials: true,
		}),
		middleware.Logger,
		middleware.Recoverer,
		httprate.LimitAll(1_000, time.Minute),
		httprate.LimitByRealIP(100, time.Minute),
	)
}

// UserRouter returns a configured router that handles all user endpoints.
func UserRouter(app *App) http.Handler {
	router := chi.NewRouter()

	router.Post("/", app.CreateUser)
	router.Get("/", app.GetAllUsers)
	router.Get("/{name}", app.GetUserByName)
	router.Put("/", app.JWT(app.UpdateUser))
	router.Delete("/", app.BasicAuth(app.DeleteUser))

	return router
}

// PostRouter returns a configured router that handles all post endpoints.
func PostRouter(app *App) http.Handler {
	router := chi.NewRouter()

	router.Post("/", app.JWT(app.CreatePost))
	router.Get("/", app.JWT(app.GetAllPosts))
	router.Put("/{id}", app.JWT(app.UpdatePost))
	router.Delete("/{id}", app.JWT(app.DeletePost))

	return router
}

func PostLikeRouter(app *App) http.Handler {
	router := chi.NewRouter()

	router.Post("/", app.JWT(app.LikePost))
	router.Delete("/", app.JWT(app.UnlikePost))

	return router
}

func AdminRouter(app *App) http.Handler {
	router := chi.NewRouter()

	router.Post("/jwt", app.JWT(app.ValidateJWT))
	router.Get("/error", HandleError)
	router.Post("/panic", HandlePanic)
	router.Post("/sleep", HandleSleep)

	return router
}
