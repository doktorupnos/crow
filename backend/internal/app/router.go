package app

import (
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/env"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func ConfiguredRouter(env *env.Env) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{env.CorsOrigin},
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
	router.Mount("/admin", AdminRouter())

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
