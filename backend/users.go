package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Name     string `json:"name" gorm:"size:20;unique;not null"`
	Password string `json:"-"    gorm:"size:64;not null"`
}

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

	user := User{
		Name:     body.Name,
		Password: string(hashed),
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
