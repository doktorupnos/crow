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
	ErrWrongPassword    = AuthError("wrong password")
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User) error

func (s *State) BasicAuth(handler AuthHandler) ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		username, password, ok := r.BasicAuth()
		if !ok {
			return APIError{
				Code: http.StatusBadRequest,
				Err:  ErrMissingBasicAuth,
			}
		}

		user, err := s.DB.GetUserByName(r.Context(), username)
		if err != nil {
			return APIError{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		if !PasswordsMatch(user.Password, password) {
			return APIError{
				Code: http.StatusUnauthorized,
				Err:  ErrWrongPassword,
			}
		}

		return handler(w, r, user)
	}
}

func (s *State) Login(w http.ResponseWriter, r *http.Request, user database.User) error {
	respond.JWT(w, http.StatusOK, s.Secret, user.ID.String(), s.ExpiresIn)
	return nil
}

func (s *State) Logout(w http.ResponseWriter, r *http.Request, user database.User) error {
	respond.JWT(w, http.StatusOK, s.Secret, user.ID.String(), 0)
	return nil
}

func (s *State) JWT(handler AuthHandler) ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		c, err := r.Cookie("token")
		if err != nil {
			return APIError{
				Code: http.StatusUnauthorized,
				Err:  err,
			}
		}

		token, err := jwt.Parse(s.Secret, c.Value)
		if err != nil {
			return APIError{
				Code: http.StatusUnauthorized,
				Err:  err,
			}
		}

		subject, err := token.Claims.GetSubject()
		if err != nil {
			return APIError{
				Code: http.StatusUnauthorized,
				Err:  err,
			}
		}

		userID, err := uuid.Parse(subject)
		if err != nil {
			return APIError{
				Code: http.StatusUnauthorized,
				Err:  err,
			}
		}

		user, err := s.DB.GetUserByID(r.Context(), userID)
		if err != nil {
			return APIError{
				Code: http.StatusUnauthorized,
				Err:  err,
			}
		}

		return handler(w, r, user)
	}
}

func (s *State) ValidateJWT(w http.ResponseWriter, _ *http.Request, _ database.User) error {
	w.WriteHeader(http.StatusOK)
	return nil
}
