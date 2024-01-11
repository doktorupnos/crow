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
	Following      bool      `json:"following"`
}

func (app *App) ViewProfile(w http.ResponseWriter, r *http.Request, u user.User) {
	var err error
	var self, following bool

	target := u

	name := r.URL.Query().Get("u")
	if name != "" {
		target, err = app.userService.GetByName(name)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	if target.Name == u.Name {
		self = true
	} else {
		// set following if the user 'u' follows the target
		following, err = app.userService.FollowsUser(u, target)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	followingCount, err := app.userService.FollowingCount(target)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	followerCount, err := app.userService.FollowerCount(target)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := &ProfileResponse{
		Self:           self,
		UserID:         target.ID,
		UserName:       target.Name,
		FollowingCount: followingCount,
		FollowerCount:  followerCount,
		Following:      following,
	}
	respondWithJSON(w, http.StatusOK, resp)
}
