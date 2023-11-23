package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// registerRoutes groups the definition of all the endpoints.
func registerRoutes(app *App) http.Handler {
	mainRouter := chi.NewRouter()

	// NOTE: cors not final
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
	// NOTE: Explore chi's middleware package for more userful handlers
	mainRouter.Use(middleware.Logger)

	// health-check, readiness endpoint.
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

	mainRouter.Post("/login", app.WithBasicAuth(app.Login))

	userRouter := chi.NewRouter()
	userRouter.Route("/", func(r chi.Router) {
		r.Post("/", app.CreateUser)
		r.Get("/{name}", app.GetUserByName)
		r.Get("/", app.GetAllUsers)
		r.Put("/", app.WithJWT(app.UpdateUser))
		r.Delete("/", app.WithBasicAuth(app.DeleteUser))
	})
	mainRouter.Mount("/users", userRouter)

	return mainRouter
}
