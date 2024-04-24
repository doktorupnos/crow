package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func ConfiguredRouter(app *App) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
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
	}))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/healthz", HealthCheck)

	router.Post("/login", app.BasicAuth(app.Login))
	router.Post("/logout", app.JWT(app.Logout))

	router.Mount("/users", UserRouter(app))
	router.Get("/profile", app.JWT(app.ViewProfile))

	router.Post("/follow", app.JWT(app.Follow))
	router.Post("/unfollow", app.JWT(app.UnFollow))
	router.Get("/following", app.JWT(app.Following))
	router.Get("/followers", app.JWT(app.Followers))
	router.Get("/following_count", app.JWT(app.FollowingCount))
	router.Get("/followers_count", app.JWT(app.FollowerCount))

	router.Mount("/posts", PostRouter(app))
	router.Mount("/post_likes", PostLikeRouter(app))

	router.Mount("/admin", AdminRouter(app))

	return router
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
