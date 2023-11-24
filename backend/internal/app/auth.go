package app

import (
	"fmt"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/google/uuid"
)

type AuthenticatedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

// WithBasicAuth authenticates a User using basic username:password Authorization.
func (app *App) WithBasicAuth(handler AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name, password, ok := r.BasicAuth()
		if !ok {
			respondWithError(w, http.StatusUnauthorized, "malformed Authorization header")
			return
		}

		user, err := database.GetUserByName(app.DB, name)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		if !PasswordsMatch(user.Password, password) {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("passwords do not match"))
			return
		}

		handler(w, r, user)
	}
}

// WithJWT authenticates a user's JWT token through cookies.
func (app *App) WithJWT(handler AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		tokenString := c.Value

		token, err := ParseJWT(app.JWT_SECRET, tokenString)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		idString, err := token.Claims.GetSubject()
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		id, err := uuid.Parse(idString)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := database.GetUserByID(app.DB, id)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		handler(w, r, user)
	}
}
