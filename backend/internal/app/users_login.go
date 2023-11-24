package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close()

	signedToken, err := NewJWT(app.JWT_SECRET, user.ID.String(), app.JWT_EXPIRES_IN_MINUTES)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: signedToken,
	})
}
