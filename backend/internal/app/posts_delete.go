package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *App) DeletePost(w http.ResponseWriter, r *http.Request, user database.User) {
	postIDStr := chi.URLParam(r, "id")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = database.DeletePost(app.DB, postID, user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
