package app

import (
	"encoding/json"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/respond"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

func (app *App) Follow(w http.ResponseWriter, r *http.Request, u user.User) {
	type RequestBody struct {
		Id string `json:"user_id"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	id, err := uuid.Parse(body.Id)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := app.followService.Follow(u, id); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *App) UnFollow(w http.ResponseWriter, r *http.Request, u user.User) {
	type RequestBody struct {
		Id string `json:"user_id"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	id, err := uuid.Parse(body.Id)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := app.followService.Unfollow(u, id); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *App) Following(w http.ResponseWriter, r *http.Request, u user.User) {
	target, err := app.extractTargetUser(r, u)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	us, err := app.followService.Following(r, app.Env.Pagination.DefaultFollowPageSize, target.ID)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusOK, us)
}

func (app *App) Followers(w http.ResponseWriter, r *http.Request, u user.User) {
	target, err := app.extractTargetUser(r, u)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	us, err := app.followService.Followers(r, app.Env.Pagination.DefaultFollowPageSize, target.ID)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusOK, us)
}

func (app *App) FollowingCount(w http.ResponseWriter, r *http.Request, u user.User) {
	target, err := app.extractTargetUser(r, u)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	count, err := app.followService.FollowingCount(target)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusOK, count)
}

func (app *App) FollowerCount(w http.ResponseWriter, r *http.Request, u user.User) {
	target, err := app.extractTargetUser(r, u)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	count, err := app.followService.FollowerCount(target)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusOK, count)
}
