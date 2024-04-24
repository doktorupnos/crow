package app

import (
	"net/http"
	"time"

	gracefulshutdown "github.com/quii/go-graceful-shutdown"
)

func GracefulServer(app *App, router http.Handler) *gracefulshutdown.Server {
	server := &http.Server{
		Addr:    app.Env.Server.Addr,
		Handler: router,

		// Good practice to set timeouts to avoid Slowloris attacks.
		// Values are not final, should be extracted into configuration variables.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Minute,
	}
	return gracefulshutdown.NewServer(server)
}
