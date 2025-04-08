package app

import (
	"encoding/json"
	"net/http"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/doktorupnos/crow/api/internal/respond"
	"github.com/google/uuid"
)

type FollowRequest struct {
	UserID string `json:"user_id"`
}

func (s *State) CreateFollow(w http.ResponseWriter, r *http.Request, user database.User) error {
	var req FollowRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	targetID, err := uuid.Parse(req.UserID)
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	err = s.db.CreateFollow(r.Context(), database.CreateFollowParams{
		Follower: user.ID,
		Followee: targetID,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (s *State) DeleteFollow(w http.ResponseWriter, r *http.Request, user database.User) error {
	var req FollowRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	targetID, err := uuid.Parse(req.UserID)
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	err = s.db.DeleteFollow(r.Context(), database.DeleteFollowParams{
		Follower: user.ID,
		Followee: targetID,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

// TODO: same code. Only method changes. Restructure later

func (s *State) GetFollowerCount(w http.ResponseWriter, r *http.Request, user database.User) error {
	target := user.ID

	if r.URL.Query().Has("u") {
		id, err := uuid.Parse(r.URL.Query().Get("u"))
		if err != nil {
			return APIError{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		target = id
	}

	count, err := s.db.GetFollowerCount(r.Context(), target)
	if err != nil {
		return err
	}

	respond.JSON(w, http.StatusOK, count)
	return nil
}

func (s *State) GetFollowingCount(w http.ResponseWriter, r *http.Request, user database.User) error {
	target := user.ID

	if r.URL.Query().Has("u") {
		id, err := uuid.Parse(r.URL.Query().Get("u"))
		if err != nil {
			return APIError{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		target = id
	}

	count, err := s.db.GetFollowingCount(r.Context(), target)
	if err != nil {
		return err
	}

	respond.JSON(w, http.StatusOK, count)
	return nil
}

func (s *State) GetFollowers(w http.ResponseWriter, r *http.Request, user database.User) error {
	target := user

	queries := r.URL.Query()
	key := "u"
	if queries.Has(key) {
		name := queries.Get(key)
		u, err := s.db.GetUserByName(r.Context(), name)
		if err != nil {
			return APIError{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}
		target = u
	}

	pages, err := extractPagination(r)
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	followers, err := s.db.GetFollowers(r.Context(), database.GetFollowersParams{
		Followee: target.ID,
		Limit:    pages.Limit,
		Offset:   pages.Offset,
	})
	if err != nil {
		return err
	}

	respond.JSON(w, http.StatusOK, followers)
	return nil
}

func (s *State) GetFollowing(w http.ResponseWriter, r *http.Request, user database.User) error {
	target := user

	queries := r.URL.Query()
	key := "u"
	if queries.Has(key) {
		name := queries.Get(key)
		u, err := s.db.GetUserByName(r.Context(), name)
		if err != nil {
			return APIError{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}
		target = u
	}

	pages, err := extractPagination(r)
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	following, err := s.db.GetFollowing(r.Context(), database.GetFollowingParams{
		Follower: target.ID,
		Limit:    pages.Limit,
		Offset:   pages.Offset,
	})
	if err != nil {
		return err
	}

	respond.JSON(w, http.StatusOK, following)
	return nil
}
