package main

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/app"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// registerRoutes groups the definition of all the endpoints.
func registerRoutes(app *app.App) http.Handler {
	mainRouter := chi.NewRouter()

	// NOTE: cors not final
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{app.ORIGIN},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	// NOTE: Explore chi's middleware package for more userful handlers
	mainRouter.Use(middleware.Logger)

	// health-check, readiness endpoint.
	mainRouter.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		// TODO: enhance to return a http.StatusServiceUnavailable.
		w.WriteHeader(http.StatusOK)
	})

	mainRouter.Post("/login", app.WithBasicAuth(app.Login))
	mainRouter.Post("/logout", app.WithJWT(app.Logout))

	usersRouter := chi.NewRouter()
	usersRouter.Route("/", func(r chi.Router) {
		r.Post("/", app.CreateUser)
		r.Get("/{name}", app.GetUserByName)
		r.Get("/", app.GetAllUsers)
		r.Put("/", app.WithJWT(app.UpdateUser))
		r.Delete("/", app.WithBasicAuth(app.DeleteUser))
	})
	mainRouter.Mount("/users", usersRouter)

	postsRouter := chi.NewRouter()
	postsRouter.Route("/", func(r chi.Router) {
		r.Post("/", app.WithJWT(app.CreatePost))
		r.Get("/", app.WithJWT(app.RetrievePosts))
		r.Put("/{id}", app.WithJWT(app.UpdatePost))
		r.Delete("/{id}", app.WithJWT(app.DeletePost))
	})
	mainRouter.Mount("/posts", postsRouter)

	return mainRouter
}
