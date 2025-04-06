package app

import (
	"encoding/json"
	"net/http"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/google/uuid"
)

type LikeRequest struct {
	PostID uuid.UUID `json:"post_id"`
}

func (s *State) CreateLike(w http.ResponseWriter, r *http.Request, user database.User) error {
	var req LikeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	err := s.DB.CreateLike(r.Context(), database.CreateLikeParams{
		UserID: user.ID,
		PostID: req.PostID,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}

func (s *State) DeleteLike(w http.ResponseWriter, r *http.Request, user database.User) error {
	var req LikeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	err := s.DB.DeleteLike(r.Context(), database.DeleteLikeParams{
		UserID: user.ID,
		PostID: req.PostID,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
