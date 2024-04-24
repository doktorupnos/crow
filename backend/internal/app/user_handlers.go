package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/doktorupnos/crow/backend/internal/respond"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var errDecodeRequestBody = errors.New("failed to decode request body")

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, errDecodeRequestBody)
		return
	}

	userID, err := app.userService.Create(body.Name, body.Password)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			err = ErrUserNameTaken
		}
		respond.Error(w, http.StatusBadRequest, err)
		return
	}
	respond.JWT(w, http.StatusCreated, app.Env.JwtSecret, userID.String(), app.Env.JwtLifetime)
}

func (app *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.userService.GetAll()
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusOK, users)
}

func (app *App) GetUserByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		respond.Error(w, http.StatusBadRequest, errMissingURLParameter)
		return
	}

	u, err := app.userService.GetByName(name)
	if err != nil {
		respond.Error(w, http.StatusNotFound, err)
		return
	}
	respond.JSON(w, http.StatusOK, u)
}

func (app *App) DeleteUser(w http.ResponseWriter, r *http.Request, u user.User) {
	if err := app.userService.Delete(u); err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *App) UpdateUser(w http.ResponseWriter, r *http.Request, u user.User) {
	type RequestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, errDecodeRequestBody)
		return
	}

	if err := app.userService.Update(u, body.Name, body.Password); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	u, err = app.userService.GetByID(u.ID)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
	}
	respond.JSON(w, http.StatusOK, u)
}

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

	if err := app.userService.Follow(u, id); err != nil {
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

	if err := app.userService.Unfollow(u, id); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (app *App) extractTargetUser(r *http.Request, u user.User) (user.User, error) {
	target := u
	if uq := r.URL.Query().Get("u"); uq != "" {
		var err error
		target, err = app.userService.GetByName(uq)
		if err != nil {
			return user.User{}, err
		}
	}
	return target, nil
}

func (app *App) Following(w http.ResponseWriter, r *http.Request, u user.User) {
	target, err := app.extractTargetUser(r, u)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	us, err := app.userService.Following(r, app.Env.DefaultFollowPageSize, target.ID)
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

	us, err := app.userService.Followers(r, app.Env.DefaultFollowPageSize, target.ID)
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

	count, err := app.userService.FollowingCount(target)
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

	count, err := app.userService.FollowerCount(target)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusOK, count)
}
