package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
)

func (app *App) UpdateUser(w http.ResponseWriter, r *http.Request, user database.User) {
	type RequestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("failed to parse request body : %s", err),
		)
		return
	}

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("failed to hash passwrod : %s", err))
		return
	}

	_, err = database.UpdateUser(app.DB, user.ID, database.UpdateUserParams{
		Name:     body.Name,
		Password: hashedPassword,
	})
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("failed to update user : %s", err),
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}
