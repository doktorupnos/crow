package app

import (
	"net/http"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/doktorupnos/crow/api/internal/jwt"
	"github.com/doktorupnos/crow/api/internal/respond"
	"github.com/google/uuid"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

const (
	ErrMissingBasicAuth = AuthError("missing Authorizaton Basic header")
	ErrPassword         = AuthError("wrong password")
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func (s *State) BasicAuth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			respond.Error(w, http.StatusBadRequest, ErrMissingBasicAuth)
			return
		}

		user, err := s.DB.GetUserByName(r.Context(), username)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, err)
			return
		}

		if !PasswordsMatch(user.Password, password) {
			respond.Error(w, http.StatusUnauthorized, ErrPassword)
			return
		}

		handler(w, r, user)
	}
}

func (s *State) Login(w http.ResponseWriter, r *http.Request, user database.User) {
	respond.JWT(w, http.StatusOK, s.Secret, user.ID.String(), s.ExpiresIn)
}

func (s *State) Logout(w http.ResponseWriter, r *http.Request, user database.User) {
	respond.JWT(w, http.StatusOK, s.Secret, user.ID.String(), 0)
}

func (s *State) JWT(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		token, err := jwt.Parse(s.Secret, c.Value)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		subject, err := token.Claims.GetSubject()
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		userID, err := uuid.Parse(subject)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		user, err := s.DB.GetUserByID(r.Context(), userID)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		handler(w, r, user)
	}
}

func (s *State) ValidateJWT(w http.ResponseWriter, _ *http.Request, _ database.User) {
	w.WriteHeader(http.StatusOK)
}
