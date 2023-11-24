package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
)

func (app *App) CreatePost(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close()

	type ResponseBody struct {
		Body string `json:"body"`
	}
	body := ResponseBody{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateBody(body.Body); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := database.CreatePost(app.DB, database.CreatePostParams{
		Body:   body.Body,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, post)
}

func validateBody(body string) error {
	if len(body) == 0 {
		return errors.New("empty body")
	}

	return nil
}
