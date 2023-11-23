package main

import (
	"fmt"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/go-chi/chi/v5"
)

func (app *App) GetUserByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		respondWithError(w, http.StatusBadRequest, "missing {name} url parameter")
		return
	}

	user, err := database.GetUserByName(app.DB, name)
	if err != nil {
		respondWithError(
			w,
			http.StatusNotFound,
			fmt.Sprintf("failed to retrieve user : %s", err.Error()),
		)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (app *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers(app.DB)
	if err != nil {
		respondWithError(
			w,
			http.StatusNotFound,
			fmt.Sprintf("failed to retrieve users : %s", err.Error()),
		)
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}
