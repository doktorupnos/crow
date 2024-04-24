package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/jwt"
	"github.com/doktorupnos/crow/backend/internal/passwd"
	"github.com/doktorupnos/crow/backend/internal/respond"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

const (
	ErrMissingAuthBasicHeader = AuthError("missing Authorization Basic header")
	ErrWrongPassword          = AuthError("wrong password")
)

type authedHandler func(w http.ResponseWriter, r *http.Request, u user.User)

func (app *App) BasicAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			respond.Error(w, http.StatusUnauthorized, ErrMissingAuthBasicHeader)
			return
		}

		u, err := app.userService.GetByName(username)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		if !passwd.Match(u.Password, password) {
			respond.Error(w, http.StatusUnauthorized, ErrWrongPassword)
			return
		}

		handler(w, r, u)
	}
}

func (app *App) JWT(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		tokenString := c.Value

		token, err := jwt.Parse(app.Env.JWT.Secret, tokenString)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		userIDString, err := token.Claims.GetSubject()
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		userID, err := uuid.Parse(userIDString)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		u, err := app.userService.GetByID(userID)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		handler(w, r, u)
	}
}

func (app *App) ValidateJWT(w http.ResponseWriter, r *http.Request, u user.User) {
	w.WriteHeader(http.StatusOK)
}
