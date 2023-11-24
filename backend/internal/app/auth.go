package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/doktorupnos/crow/backend/internal/database"
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

// WithJWT authenticates a user's JWT token.
func (app *App) WithJWT(handler AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			respondWithJSON(w, http.StatusUnauthorized, "missing Authorization header")
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) != 2 {
			respondWithError(w, http.StatusUnauthorized, "malformed Authorization header")
			return
		}

		const methodBearer = "Bearer"
		authMethod := fields[0]
		if authMethod != methodBearer {
			respondWithError(w, http.StatusUnauthorized, "unsupported Authorization method")
			return
		}

		// TODO: Improve Error Messages

		tokenString := fields[1]
		token, err := ParseJWT(app.JWT_SECRET, tokenString)
		if err != nil {
			respondWithError(
				w,
				http.StatusUnauthorized,
				fmt.Sprintf("failed to parse token : %s", err),
			)
			return
		}

		username, err := token.Claims.GetSubject()
		if err != nil {
			respondWithError(
				w,
				http.StatusUnauthorized,
				fmt.Sprintf("failed to parse token : %s", err),
			)
			return
		}

		user, err := database.GetUserByName(app.DB, username)
		if err != nil {
			respondWithError(
				w,
				http.StatusUnauthorized,
				fmt.Sprintf("failed to retrieve user : %s", err),
			)
			return
		}

		handler(w, r, user)
	}
}
