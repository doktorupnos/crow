package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/respond"
	"github.com/doktorupnos/crow/backend/internal/user"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request, u user.User) {
	respond.JWT(w, http.StatusOK, app.Env.JwtSecret, u.ID.String(), app.Env.JwtLifetime)
}

func (app *App) Logout(w http.ResponseWriter, r *http.Request, u user.User) {
	respond.JWT(w, http.StatusOK, app.Env.JwtSecret, u.ID.String(), 0)
}
