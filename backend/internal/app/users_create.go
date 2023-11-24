package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
)

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type RequestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	body := RequestBody{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&body)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("failed to decode request body : %s", err),
		)
		return
	}

	if err := validateName(body.Name); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validatePassword(body.Password); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("failed to hash password : %s", err.Error()),
		)
		return
	}

	user, err := database.CreateUser(app.DB, database.CreateUserParams{
		Name:     body.Name,
		Password: hashedPassword,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("failed to create user : %s", err.Error()),
		)
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func validateName(name string) error {
	// TODO: Validate username to only contain alphabetic characters, numbers, and the underscore.

	if len(name) == 0 {
		return errors.New("empty name")
	}

	return nil
}

func validatePassword(password string) error {
	// TODO: Validate password length for bcrypt's maximum length of 72 bytes
	if len(password) == 0 {
		return errors.New("empty password")
	}

	return nil
}
