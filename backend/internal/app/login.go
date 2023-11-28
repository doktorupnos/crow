package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/jwt"
	"github.com/doktorupnos/crow/backend/internal/user"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request, u user.User) {
	signedToken, err := jwt.Create(app.Env.JwtSecret, u.ID.String(), app.Env.JwtLifetime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: signedToken,
	})
	w.WriteHeader(http.StatusOK)
}

func (app *App) Logout(w http.ResponseWriter, r *http.Request, u user.User) {
	signedToken, err := jwt.Create(app.Env.JwtSecret, u.ID.String(), 0)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: signedToken,
	})
	w.WriteHeader(http.StatusOK)
}
