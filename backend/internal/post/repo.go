package post

import "github.com/google/uuid"

type Repo interface {
	Create(p Post) error
	Load(params LoadParams) ([]FeedPost, error)
	LoadAllByID(params LoadParams) ([]FeedPost, error)
	LoadByID(id uuid.UUID) (Post, error)
	Update(p Post) error
	Delete(p Post) error
}
