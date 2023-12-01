package app

import (
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

func (s *PostService) Create(u user.User, body string) error {
	if err := validateBody(body); err != nil {
		return err
	}

	p := post.Post{Body: body, UserID: u.ID, User: u}
	return s.pr.Create(p)
}

func (s *PostService) Load(params post.PaginationParams) ([]post.FeedPost, error) {
	return s.pr.Load(params)
}

func (s *PostService) LoadByID(id uuid.UUID) (post.Post, error) {
	return s.pr.LoadByID(id)
}

func (s *PostService) Update(postID, userID uuid.UUID, body string) error {
	if err := validateBody(body); err != nil {
		return err
	}

	p, err := s.LoadByID(postID)
	if err != nil {
		return err
	}

	if p.UserID != userID {
		return PostErrNotOwner
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
		return PostErrNotOwner
	}

	return s.pr.Delete(p)
}

type PostErr string

func (e PostErr) Error() string {
	return string(e)
}

const (
	PostErrEmpty    = PostErr("Post is empty")
	PostErrTooBig   = PostErr("Post is too big")
	PostErrNotOwner = PostErr("Post does not belong this user")
)

func validateBody(body string) error {
	if len(body) == 0 {
		return PostErrEmpty
	}

	if len(body) > 128 {
		return PostErrTooBig
	}

	return nil
}
