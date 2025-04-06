package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/doktorupnos/crow/api/internal/respond"
	"github.com/google/uuid"
)

type UserServer struct {
	Service      UserService
	JWTSecret    string
	JWTExpiresIn time.Duration
}

type UserService interface {
	Create(context.Context, CreateUserRequest) (database.User, error)
}

type UserServicePG struct {
	db *database.Queries
}

var _ UserService = (*UserServicePG)(nil)

func (s *UserServicePG) Create(ctx context.Context, r CreateUserRequest) (database.User, error) {
	if err := r.Validate(); err != nil {
		return database.User{}, APIError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid request body: %w", err),
		}
	}

	hashedPassword, err := Hash(r.Password)
	if err != nil {
		return database.User{}, APIError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid password: %w", err),
		}
	}

	now := time.Now()
	user, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      r.Name,
		Password:  hashedPassword,
	})
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return database.User{}, APIError{
				Code: http.StatusBadRequest,
				Err:  ErrUsernameTaken,
			}
		}
		return database.User{}, err
	}

	return user, nil
}

func (s *UserServer) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var req CreateUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("invalid JSON: %w", err),
		}
	}

	user, err := s.Service.Create(r.Context(), req)
	if err != nil {
		return err
	}

	// TODO: replace with slog
	log.Printf("%#v\n", user)

	respond.JWT(
		w,
		http.StatusCreated,
		s.JWTSecret,
		user.ID.String(),
		s.JWTExpiresIn,
	)
	return nil
}

type CreateUserError string

func (e CreateUserError) Error() string {
	return string(e)
}

const (
	ErrUsernameEmpty     CreateUserError = "username is empty"
	ErrUsernameSmall     CreateUserError = "username is too small"
	ErrUsernameBig       CreateUserError = "username is too big"
	ErrUsernameMalformed CreateUserError = "malformed user name"
	ErrUsernameTaken     CreateUserError = "username taken"

	ErrPasswordEmpty     CreateUserError = "password is empty"
	ErrPasswordBig       CreateUserError = "password is too big"
	ErrPasswordMalformed CreateUserError = "malformed password"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r CreateUserRequest) Validate() error {
	if err := validateUsername(r.Name); err != nil {
		return err
	}
	if err := validatePassword(r.Password); err != nil {
		return err
	}
	return nil
}

var (
	usernameRegex = regexp.MustCompile("^[a-zA-Z0-9_.]+$")
	passwordRegex = regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*]")
)

func validateUsername(name string) error {
	switch l := len(name); {
	case l == 0:
		return ErrUsernameEmpty
	case l < 4:
		return ErrUsernameSmall
	case l > 20:
		return ErrUsernameBig
	}
	if !usernameRegex.MatchString(name) {
		return ErrUsernameMalformed
	}
	return nil
}

func validatePassword(password string) error {
	switch l := len(password); {
	case l == 0:
		return ErrPasswordEmpty
	case l > 72:
		return ErrPasswordBig
	}
	if !passwordRegex.MatchString(password) {
		return ErrPasswordMalformed
	}
	return nil
}
