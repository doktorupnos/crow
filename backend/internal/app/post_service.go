package app

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type PostService struct {
	pr post.PostRepo
}

func NewPostService(pr post.PostRepo) *PostService {
	return &PostService{pr}
}

func (s *PostService) Create(u user.User, body string, limit int) error {
	if err := validateBody(body, limit); err != nil {
		return err
	}

	p := post.Post{Body: body, UserID: u.ID, User: u}
	return s.pr.Create(p)
}

func (s *PostService) Load(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]post.FeedPost, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}
	order := "desc"

	return s.pr.Load(post.LoadParams{
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
		Order:  order,
		UserID: userID,
	})
}

func (s *PostService) LoadAllByID(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]post.FeedPost, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}
	order := "desc"

	return s.pr.LoadAllByID(post.LoadParams{
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
		Order:  order,
		UserID: userID,
	})
}

func (s *PostService) LoadByID(id uuid.UUID) (post.Post, error) {
	return s.pr.LoadByID(id)
}

func (s *PostService) Update(postID, userID uuid.UUID, body string, limit int) error {
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
	return s.pr.Update(p)
}

func (s *PostService) Delete(id, userID uuid.UUID) error {
	p, err := s.LoadByID(id)
	if err != nil {
		return err
	}

	if p.UserID != userID {
		return ErrNotPostOwner
	}

	return s.pr.Delete(p)
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
