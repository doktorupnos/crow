package app

import (
	"github.com/doktorupnos/crow/backend/internal/like"
	"github.com/google/uuid"
)

type PostLikeService struct {
	postLikeRepo like.PostLikeRepo
}

func NewPostLikeService(lr like.PostLikeRepo) *PostLikeService {
	return &PostLikeService{lr}
}

func (s *PostLikeService) Create(userID, postID uuid.UUID) error {
	return s.postLikeRepo.Create(like.PostLike{UserID: userID, PostID: postID})
}

func (s *PostLikeService) Load() ([]like.PostLike, error) {
	return s.postLikeRepo.Load()
}

func (s *PostLikeService) Single(userID, postID uuid.UUID) (like.PostLike, error) {
	return s.postLikeRepo.Single(userID, postID)
}

func (s *PostLikeService) Delete(userID, postID uuid.UUID) error {
	l, err := s.postLikeRepo.Single(userID, postID)
	if err != nil {
		return err
	}
	return s.postLikeRepo.Delete(l)
}
