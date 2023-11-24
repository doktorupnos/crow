package app

import (
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/database"
)

func (app *App) Logout(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close()

	signedToken, err := NewJWT(app.JWT_SECRET, user.Name, 0)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   signedToken,
		Expires: time.Now(), // NOTE: probably don't need this
	})
}
