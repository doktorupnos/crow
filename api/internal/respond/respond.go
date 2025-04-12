package respond

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/api/internal/jwt"
)

func JSON(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(data); err != nil {
		slog.Error("could not write JSON response", "error", err.Error())
		http.Error(w, "could not write JSON response", http.StatusInternalServerError)
	}
}

type ErrorResponse struct {
	Message string `json:"error"`
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, ErrorResponse{err.Error()})
}

func JWT(
	w http.ResponseWriter,
	statusCode int,
	secret, userID string,
	expiresIn time.Duration,
) {
	signed, err := jwt.Create(secret, userID, expiresIn)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Path:  "/",
		Value: signed,
	})
	w.WriteHeader(statusCode)
}
