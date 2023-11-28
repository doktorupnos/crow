package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/user"
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
	posts, err := app.postService.GetAll()
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	type ResponseItem struct {
		ID        uuid.UUID `json:"id"`
		UserID    uuid.UUID `json:"user_id"`
		UserName  string    `json:"user_name"`
		CreatedAt time.Time `json:"created_at"`
		Body      string    `json:"body"`
	}

	response := make([]ResponseItem, 0, len(posts))
	for _, post := range posts {
		response = append(response, ResponseItem{
			ID:        post.ID,
			UserID:    post.UserID,
			UserName:  post.User.Name,
			CreatedAt: post.CreatedAt,
			Body:      post.Body,
		})
	}

	respondWithJSON(w, http.StatusOK, response)
}
