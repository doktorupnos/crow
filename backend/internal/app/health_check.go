package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/respond"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(
		w,
		http.StatusOK,
		HealthCheckResponse{http.StatusText(http.StatusOK)},
	)
}
