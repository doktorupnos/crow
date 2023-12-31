package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/jwt"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, statusCode, errorResponse{message})
}

func respondWithJWT(
	w http.ResponseWriter,
	statusCode int,
	secret, subject string,
	lifetime time.Duration,
) {
	signedToken, err := jwt.Create(secret, subject, lifetime)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: signedToken,
	})
	w.WriteHeader(statusCode)
}
