package app

import (
	"net/http"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/doktorupnos/crow/api/internal/respond"
	"github.com/google/uuid"
)

type ProfileResponse struct {
	UserID    uuid.UUID `json:"id"`
	UserName  string    `json:"name"`
	Followers int64     `json:"follower_count"`
	Following int64     `json:"following_count"`
	Self      bool      `json:"self"`
	Follows   bool      `json:"following"`
}

func (s *State) Profile(w http.ResponseWriter, r *http.Request, user database.User) error {
	queries := r.URL.Query()

	target := user
	if queries.Has("u") {
		name := queries.Get("u")
		u, err := s.db.GetUserByName(r.Context(), name)
		if err != nil {
			return APIError{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}
		target = u
	}

	self := target.ID == user.ID

	followers, err := s.db.GetFollowerCount(r.Context(), target.ID)
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	following, err := s.db.GetFollowingCount(r.Context(), target.ID)
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	follows, err := s.db.FollowsUser(r.Context(), database.FollowsUserParams{
		Follower: user.ID,
		Followee: target.ID,
	})
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	resp := ProfileResponse{
		UserID:    target.ID,
		UserName:  target.Name,
		Followers: followers,
		Following: following,
		Self:      self,
		Follows:   follows,
	}
	respond.JSON(w, http.StatusOK, resp)

	return nil
}
