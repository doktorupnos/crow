package follow

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type Service struct {
	r Repo
	s *user.Service
}

func NewService(r Repo, s *user.Service) *Service {
	return &Service{r, s}
}

func (s *Service) Follow(u user.User, id uuid.UUID) error {
	o, err := s.s.GetByID(id)
	if err != nil {
		return err
	}

	return s.r.Follow(u, o)
}

func (s *Service) Unfollow(u user.User, id uuid.UUID) error {
	o, err := s.s.GetByID(id)
	if err != nil {
		return err
	}

	return s.r.Unfollow(u, o)
}

func (s *Service) Following(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]Follow, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}

	return s.r.Following(LoadParams{
		UserID: userID,
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
	})
}

func (s *Service) Followers(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]Follow, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}

	return s.r.Followers(LoadParams{
		UserID: userID,
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
	})
}

func (s *Service) FollowingCount(u user.User) (int, error) {
	return s.r.FollowingCount(u)
}

func (s *Service) FollowerCount(u user.User) (int, error) {
	return s.r.FollowersCount(u)
}

func (s *Service) FollowsUser(u, t user.User) (bool, error) {
	return s.r.FollowsUser(u, t)
}
