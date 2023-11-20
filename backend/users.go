package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Model
	Name     string `gorm:"size:20;unique;not null"`
	Password string `gorm:"size:64;not null"`
}

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	body := RequestBody{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to parse request body")
		return
	}

	// TODO: Validate username

	user := User{
		Name:     body.Name,
		Password: body.Password,
	}

	if err := cfg.DB.Create(&user).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (cfg *ApiConfig) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	if err := cfg.DB.Find(&users).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to retrieve users")
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}
