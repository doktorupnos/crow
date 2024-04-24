package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/respond"
	"github.com/doktorupnos/crow/backend/internal/user"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request, u user.User) {
	respond.JWT(w, http.StatusOK, app.Env.JWT.Secret, u.ID.String(), app.Env.JWT.Lifetime)
}

func (app *App) Logout(w http.ResponseWriter, r *http.Request, u user.User) {
	respond.JWT(w, http.StatusOK, app.Env.JWT.Secret, u.ID.String(), 0)
}
