package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/doktorupnos/crow/backend/internal/respond"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var errMissingURLParameter = errors.New("missing URL parameter")

func (app *App) CreatePost(w http.ResponseWriter, r *http.Request, u user.User) {
	type RequestBody struct {
		Body string `json:"body"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := app.postService.Create(u, body.Body, app.Env.Posts.BodyLimit); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
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
		target := user.User{}
		target, err = app.userService.GetByName(name)
		if err != nil {
			respond.Error(w, http.StatusNotFound, err)
			return
		}
		posts, err = app.postService.LoadAllByID(r, app.Env.Pagination.DefaultPostsPageSize, target.ID)
		if err != nil {
			respond.Error(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		posts, err = app.postService.Load(r, app.Env.Pagination.DefaultPostsPageSize, u.ID)
		if err != nil {
			respond.Error(w, http.StatusInternalServerError, err)
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

	respond.JSON(w, http.StatusOK, respPosts)
}

func (app *App) UpdatePost(w http.ResponseWriter, r *http.Request, u user.User) {
	postIDString := chi.URLParam(r, "id")
	if postIDString == "" {
		respond.Error(w, http.StatusBadRequest, errMissingURLParameter)
		return
	}

	postID, err := uuid.Parse(postIDString)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	type RequestBody struct {
		Body string `json:"body"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := app.postService.Update(postID, u.ID, body.Body, app.Env.Posts.BodyLimit); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}
}

func (app *App) DeletePost(w http.ResponseWriter, r *http.Request, u user.User) {
	postIDString := chi.URLParam(r, "id")
	if postIDString == "" {
		respond.Error(w, http.StatusBadRequest, errMissingURLParameter)
		return
	}

	postID, err := uuid.Parse(postIDString)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := app.postService.Delete(postID, u.ID); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
