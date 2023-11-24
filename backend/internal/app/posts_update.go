package app

import (
	"encoding/json"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *App) UpdatePost(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close()

	postIDStr := chi.URLParam(r, "id")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	type RequestBody struct {
		Body string `json:"body"`
	}
	body := RequestBody{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateBody(body.Body); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := database.UpdatePost(app.DB, postID, user.ID, body.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, post)
}
