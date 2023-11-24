package app

import (
	"fmt"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/database"
)

func (app *App) DeleteUser(w http.ResponseWriter, r *http.Request, user database.User) {
	defer r.Body.Close()

	err := database.DeleteUser(app.DB, user)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("failed to delete user : %s", err.Error()),
		)
	}

	w.WriteHeader(http.StatusOK)
}
