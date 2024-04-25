// package respond provides helper functions for responding to HTTP requests
package respond

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/backend/internal/jwt"
)

// JSON responds with the JSON encoding of the given payload setting the status code in the process.
// Failing to marshal the payload will result in a 500 status code [Status Internal Server Error] and the error message will be written to the response.
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

// Error is simply a call to JSON. It will encode with message of the given error under the ErrorResponse type and write it as a JSON payload.
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, ErrorResponse{err.Error()})
}

// JWT respond with a signed JWT as a http.Cokie
func JWT(
	w http.ResponseWriter,
	statusCode int,
	secret, subject string,
	lifetime time.Duration,
) {
	signedToken, err := jwt.Create(secret, subject, lifetime)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Path:  "/",
		Value: signedToken,
	})
	w.WriteHeader(statusCode)
}
