package post

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type Service struct {
	r Repo
}

func NewService(r Repo) *Service {
	return &Service{r}
}

func (s *Service) Create(u user.User, body string, limit int) error {
	if err := validateBody(body, limit); err != nil {
		return err
	}

	p := Post{Body: body, UserID: u.ID, User: u}
	return s.r.Create(p)
}

func (s *Service) Load(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]FeedPost, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}
	order := "desc"

	return s.r.Load(LoadParams{
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
		Order:  order,
		UserID: userID,
	})
}

func (s *Service) LoadAllByID(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]FeedPost, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}
	order := "desc"

	return s.r.LoadAllByID(LoadParams{
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
		Order:  order,
		UserID: userID,
	})
}

func (s *Service) LoadByID(id uuid.UUID) (Post, error) {
	return s.r.LoadByID(id)
}

func (s *Service) Update(postID, userID uuid.UUID, body string, limit int) error {
	if err := validateBody(body, limit); err != nil {
		return err
	}

	p, err := s.LoadByID(postID)
	if err != nil {
		return err
	}

	if p.UserID != userID {
		return ErrNotPostOwner
	}

	p.Body = body
	return s.r.Update(p)
}

func (s *Service) Delete(id, userID uuid.UUID) error {
	p, err := s.LoadByID(id)
	if err != nil {
		return err
	}

	if p.UserID != userID {
		return ErrNotPostOwner
	}

	return s.r.Delete(p)
}

type PostErr string

func (e PostErr) Error() string {
	return string(e)
}

const (
	ErrPostEmpty    = PostErr("Post is empty")
	ErrPostTooBig   = PostErr("Post is too big")
	ErrNotPostOwner = PostErr("Post does not belong this user")
)

func validateBody(body string, limit int) error {
	if len(body) == 0 {
		return ErrPostEmpty
	}

	if len(body) > limit {
		return ErrPostTooBig
	}

	return nil
}
