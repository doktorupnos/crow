package app

import "net/http"

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	type statusResponse struct {
		Status string `json:"status"`
	}
	respondWithJSON(
		w,
		http.StatusOK,
		statusResponse{http.StatusText(http.StatusOK)},
	)
}
