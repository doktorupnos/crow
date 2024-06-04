package post

import "github.com/google/uuid"

type Repo interface {
	Create(Post) error
	Load(LoadParams) ([]FeedPost, error)
	LoadAllByID(LoadParams) ([]FeedPost, error)
	LoadByID(uuid.UUID) (Post, error)
	Update(Post) error
	Delete(Post) error
}
