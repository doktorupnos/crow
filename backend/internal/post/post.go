package post

import (
	"github.com/google/uuid"

	"github.com/doktorupnos/crow/backend/internal/model"
	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/user"
)

type Post struct {
	Body  string      `json:"body"    gorm:"not null"`
	Likes []user.User `json:"-" gorm:"many2many:post_likes;constraint:onDelete:CASCADE"`
	User  user.User   `json:"-"       gorm:"foreignKey:UserID; not null;constraint:onDelete:CASCADE"`
	model.Model
	UserID uuid.UUID `json:"user_id"`
}

type FeedPost struct {
	UserName string `json:"user_name"`
	Post
	Likes       int64 `json:"likes"`
	LikedByUser bool  `json:"liked_by_user"`
}

type LoadParams struct {
	Order string
	pages.PaginationParams
	UserID uuid.UUID
}

type PostRepo interface {
	Create(p Post) error
	Load(params LoadParams) ([]FeedPost, error)
	LoadAllByID(params LoadParams) ([]FeedPost, error)
	LoadByID(id uuid.UUID) (Post, error)
	Update(p Post) error
	Delete(p Post) error
}
