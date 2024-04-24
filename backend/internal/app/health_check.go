package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/respond"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	type statusResponse struct {
		Status string `json:"status"`
	}
	respond.JSON(
		w,
		http.StatusOK,
		statusResponse{http.StatusText(http.StatusOK)},
	)
}
