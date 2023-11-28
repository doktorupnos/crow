package app

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func ConfiguredRouter(app *App) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{app.Env.CorsOrigin},
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

	router.Mount("/admin", AdminRouter())

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

// AdminRouter returns a configured router that handles all admin endpoints.
func AdminRouter() http.Handler {
	router := chi.NewRouter()

	router.Post("/panic", func(w http.ResponseWriter, _ *http.Request) {
		panic("The server automatically recovers from panics")
	})

	router.Get("/error", func(w http.ResponseWriter, _ *http.Request) {
		respondWithError(
			w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	})

	router.Post("/sleep", func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(time.Minute)
		w.WriteHeader(http.StatusOK)
	})

	return router
}
