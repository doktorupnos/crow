package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/jwt"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, u user.User)

func (app *App) BasicAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			respondWithError(w, http.StatusUnauthorized, "missing Authorization Basic header")
			return
		}

		u, err := app.userService.GetByName(username)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if !passwordsMatch(u.Password, password) {
			respondWithError(w, http.StatusUnauthorized, "wrong password")
			return
		}

		handler(w, r, u)
	}
}

func (app *App) JWT(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		userID, err := uuid.Parse(userIDString)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		u, err := app.userService.GetByID(userID)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		handler(w, r, u)
	}
}
