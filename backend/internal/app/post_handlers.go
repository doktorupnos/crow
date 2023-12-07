package app

import (
	"encoding/json"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *App) CreatePost(w http.ResponseWriter, r *http.Request, u user.User) {
	type RequestBody struct {
		Body string `json:"body"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := app.postService.Create(u, body.Body); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (app *App) GetAllPosts(w http.ResponseWriter, r *http.Request, u user.User) {
	posts, err := app.postService.Load(r, app.Env.DefaultPageSize, u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}

func (app *App) UpdatePost(w http.ResponseWriter, r *http.Request, u user.User) {
	postIDString := chi.URLParam(r, "id")
	if postIDString == "" {
		respondWithError(w, http.StatusBadRequest, "missing URL parameter")
		return
	}

	postID, err := uuid.Parse(postIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	type RequestBody struct {
		Body string `json:"body"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := app.postService.Update(postID, u.ID, body.Body); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
}

func (app *App) DeletePost(w http.ResponseWriter, r *http.Request, u user.User) {
	postIDString := chi.URLParam(r, "id")
	if postIDString == "" {
		respondWithError(w, http.StatusBadRequest, "missing URL parameter")
		return
	}

	postID, err := uuid.Parse(postIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := app.postService.Delete(postID, u.ID); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
