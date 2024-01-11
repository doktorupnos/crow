package app

import (
	"encoding/json"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/post"
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
	type PostResponse struct {
		post.FeedPost
		Self bool `json:"self"`
	}

	var posts []post.FeedPost
	var err error

	name := r.URL.Query().Get("u")
	if name != "" {
		u, err = app.userService.GetByName(name)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		posts, err = app.postService.LoadAllByID(r, app.Env.DefaultPostsPageSize, u.ID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		posts, err = app.postService.Load(r, app.Env.DefaultPostsPageSize, u.ID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respPosts := make([]PostResponse, 0, len(posts))
	for _, p := range posts {
		var self bool
		if p.UserID == u.ID {
			self = true
		}
		respPosts = append(respPosts, PostResponse{
			FeedPost: p,
			Self:     self,
		})
	}

	respondWithJSON(w, http.StatusOK, respPosts)
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
