package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doktorupnos/wip-chat/backend/internal/database"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	body := RequestBody{}
	err := decoder.Decode(&body)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("failed to parse request body : %q", err),
		)
		return
	}

	// TODO: Validate username to only contain alphabetic characters, numbers, and the underscore.

	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to hash password")
		return
	}

	user, err := database.CreateUser(cfg.DB, database.CreateUserParams{
		Name:           body.Name,
		HashedPassword: string(hashed),
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (cfg *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {
	name, password, ok := r.BasicAuth()
	if !ok {
		respondWithError(w, http.StatusBadRequest, "malformed Authorization header")
		return
	}

	user, err := database.GetUserByName(cfg.DB, name)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("passwords do not match"))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (cfg *ApiConfig) GetUserByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		respondWithError(w, http.StatusBadRequest, "missing url parameter : {name}")
		return
	}

	user, err := database.GetUserByName(cfg.DB, name)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (cfg *ApiConfig) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetAllUsers(cfg.DB)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}
