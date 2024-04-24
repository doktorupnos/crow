package app

import (
	"encoding/json"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/respond"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

func (app *App) LikePost(w http.ResponseWriter, r *http.Request, u user.User) {
	type RequestBody struct {
		PostID string `json:"post_id"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	postID, err := uuid.Parse(body.PostID)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	err = app.postLikeService.Create(u.ID, postID)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *App) UnlikePost(w http.ResponseWriter, r *http.Request, u user.User) {
	type RequestBody struct {
		PostID string `json:"post_id"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	postID, err := uuid.Parse(body.PostID)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	err = app.postLikeService.Delete(u.ID, postID)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
