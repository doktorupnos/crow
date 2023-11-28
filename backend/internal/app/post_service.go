package app

import (
	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/doktorupnos/crow/backend/internal/user"
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

func (s *PostService) GetAll() ([]post.Post, error) {
	return s.pr.GetAll()
}

type PostErr string

func (e PostErr) Error() string {
	return string(e)
}

const (
	PostErrEmpty  = PostErr("Post is empty")
	PostErrTooBig = PostErr("Post is too big")
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
