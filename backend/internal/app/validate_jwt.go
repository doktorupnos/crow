package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/user"
)

func (app *App) ValidateJWT(w http.ResponseWriter, r *http.Request, u user.User) {
	w.WriteHeader(http.StatusOK)
}
