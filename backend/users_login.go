package main

import (
	"fmt"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/google/uuid"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request, user database.User) {
	signedToken, err := NewJWT(app.JWT_SECRET, user.ID.String())
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("failed to create JWT : %s", err),
		)
		return
	}

	type ResponseBody struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Token string    `json:"token"`
	}

	respondWithJSON(w, http.StatusOK, ResponseBody{
		ID:    user.ID,
		Name:  user.Name,
		Token: signedToken,
	})
}
