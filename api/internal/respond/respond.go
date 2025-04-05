package respond

import (
	"encoding/json"
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
	w.Write(data)
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
