package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
)

func (app *App) UpdateUser(w http.ResponseWriter, r *http.Request, user database.User) {
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

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("failed to hash password : %s", err))
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
