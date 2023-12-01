package app

import (
	"encoding/json"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/go-chi/chi/v5"
)

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	body := RequestBody{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "unable to decode request body")
		return
	}

	userID, err := app.userService.Create(body.Name, body.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJWT(w, http.StatusCreated, app.Env.JwtSecret, userID.String(), app.Env.JwtLifetime)
}

func (app *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.userService.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (app *App) GetUserByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		respondWithError(w, http.StatusBadRequest, "missing URL parameter")
		return
	}

	u, err := app.userService.GetByName(name)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, u)
}

func (app *App) DeleteUser(w http.ResponseWriter, r *http.Request, u user.User) {
	if err := app.userService.Delete(u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
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
		respondWithError(w, http.StatusBadRequest, "is this even running")
		return
	}

	if err := app.userService.Update(u, body.Name, body.Password); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	u, err = app.userService.GetByID(u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, u)
}
