package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type ProfileResponse struct {
	UserName       string    `json:"name"`
	FollowerCount  int       `json:"follower_count"`
	FollowingCount int       `json:"following_count"`
	UserID         uuid.UUID `json:"id"`
	Self           bool      `json:"self"`
}

func (app *App) ViewProfile(w http.ResponseWriter, r *http.Request, u user.User) {
	target := u
	var self bool

	name := r.URL.Query().Get("u")
	if name != "" {
		var err error
		target, err = app.userService.GetByName(name)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
		}
	}

	if target.Name == u.Name {
		self = true
	}

	followingCount, err := app.userService.FollowingCount(u)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	followerCount, err := app.userService.FollowerCount(u)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	resp := &ProfileResponse{
		Self:           self,
		UserID:         target.ID,
		UserName:       target.Name,
		FollowingCount: followingCount,
		FollowerCount:  followerCount,
	}
	respondWithJSON(w, http.StatusOK, resp)
}
