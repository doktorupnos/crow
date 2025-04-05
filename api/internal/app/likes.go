package app

import (
	"encoding/json"
	"net/http"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/doktorupnos/crow/api/internal/respond"
	"github.com/google/uuid"
)

type LikeRequest struct {
	PostID uuid.UUID `json:"post_id"`
}

func (s *State) CreateLike(w http.ResponseWriter, r *http.Request, user database.User) {
	var req LikeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	err := s.DB.CreateLike(r.Context(), database.CreateLikeParams{
		UserID: user.ID,
		PostID: req.PostID,
	})
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *State) DeleteLike(w http.ResponseWriter, r *http.Request, user database.User) {
	var req LikeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	err := s.DB.DeleteLike(r.Context(), database.DeleteLikeParams{
		UserID: user.ID,
		PostID: req.PostID,
	})
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
