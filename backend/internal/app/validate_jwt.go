package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/jwt"
	"github.com/google/uuid"
)

func (app *App) ValidateJWT(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	tokenString := c.Value

	token, err := jwt.Parse(app.Env.JwtSecret, tokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	_, err = uuid.Parse(userIDString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
}
