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

func (s *State) CreateFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	var req FollowRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	targetID, err := uuid.Parse(req.UserID)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	err = s.DB.CreateFollow(r.Context(), database.CreateFollowParams{
		Follower: user.ID,
		Followee: targetID,
	})
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *State) DeleteFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	var req FollowRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	targetID, err := uuid.Parse(req.UserID)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	err = s.DB.DeleteFollow(r.Context(), database.DeleteFollowParams{
		Follower: user.ID,
		Followee: targetID,
	})
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO: same code. Only method changes. Restructure later

func (s *State) GetFollowerCount(w http.ResponseWriter, r *http.Request, user database.User) {
	target := user.ID

	if r.URL.Query().Has("u") {
		id, err := uuid.Parse(r.URL.Query().Get("u"))
		if err != nil {
			respond.Error(w, http.StatusBadRequest, err)
			return
		}

		target = id
	}

	count, err := s.DB.GetFollowerCount(r.Context(), target)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.JSON(w, http.StatusOK, count)
}

func (s *State) GetFollowingCount(w http.ResponseWriter, r *http.Request, user database.User) {
	target := user.ID

	if r.URL.Query().Has("u") {
		id, err := uuid.Parse(r.URL.Query().Get("u"))
		if err != nil {
			respond.Error(w, http.StatusBadRequest, err)
			return
		}

		target = id
	}

	count, err := s.DB.GetFollowingCount(r.Context(), target)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.JSON(w, http.StatusOK, count)
}

func (s *State) GetFollowers(w http.ResponseWriter, r *http.Request, user database.User) {
	target := user

	queries := r.URL.Query()
	key := "u"
	if queries.Has(key) {
		name := queries.Get(key)
		u, err := s.DB.GetUserByName(r.Context(), name)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, err)
			return
		}
		target = u
	}

	pages, err := extractPagination(r)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	followers, err := s.DB.GetFollowers(r.Context(), database.GetFollowersParams{
		Followee: target.ID,
		Limit:    pages.Limit,
		Offset:   pages.Offset,
	})
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusOK, followers)
}

func (s *State) GetFollowing(w http.ResponseWriter, r *http.Request, user database.User) {
	target := user

	queries := r.URL.Query()
	key := "u"
	if queries.Has(key) {
		name := queries.Get(key)
		u, err := s.DB.GetUserByName(r.Context(), name)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, err)
			return
		}
		target = u
	}

	pages, err := extractPagination(r)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	following, err := s.DB.GetFollowing(r.Context(), database.GetFollowingParams{
		Follower: target.ID,
		Limit:    pages.Limit,
		Offset:   pages.Offset,
	})
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusOK, following)
}
