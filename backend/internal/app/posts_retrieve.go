package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/google/uuid"
)

func (app *App) RetrievePosts(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close()

	username := r.URL.Query().Get("user")
	postID := r.URL.Query().Get("post_id")

	if postID != "" {
		app.GetPostByID(w, r, postID)
		return
	}

	if username != "" {
		app.GetPostsByUserName(w, r, username)
		return
	}

	app.GetAllPosts(w, r)
}

func (app *App) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts(app.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}

func (app *App) GetPostByID(w http.ResponseWriter, r *http.Request, postIDStr string) {
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := database.GetPostByID(app.DB, postID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, post)
}

func (app *App) GetPostsByUserName(w http.ResponseWriter, r *http.Request, username string) {
	posts, err := database.GetAllPostsByUserName(app.DB, username)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}

func (app *App) GetPostsByUserID(w http.ResponseWriter, r *http.Request, userIDStr string) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	posts, err := database.GetAllPostsByUserID(app.DB, userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}
