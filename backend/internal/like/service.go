package like

import (
	"github.com/google/uuid"
)

type Service struct {
	r Repo
}

func NewService(r Repo) *Service {
	return &Service{r}
}

func (s *Service) Create(userID, postID uuid.UUID) error {
	return s.r.Create(Like{UserID: userID, PostID: postID})
}

func (s *Service) Load() ([]Like, error) {
	return s.r.Load()
}

func (s *Service) Single(userID, postID uuid.UUID) (Like, error) {
	return s.r.Single(userID, postID)
}

func (s *Service) Delete(userID, postID uuid.UUID) error {
	l, err := s.r.Single(userID, postID)
	if err != nil {
		return err
	}
	return s.r.Delete(l)
}
