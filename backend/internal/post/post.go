package post

import (
	"github.com/doktorupnos/crow/backend/internal/model"
	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type Post struct {
	model.Model

	Body   string    `json:"body"    gorm:"not null"`
	UserID uuid.UUID `json:"user_id"`
	User   user.User `json:"-"       gorm:"foreignKey:UserID; not null;constraint:onDelete:CASCADE"`

	Likes []user.User `json:"-" gorm:"many2many:post_likes;"`
}

type FeedPost struct {
	Post
	UserName    string `json:"user_name"`
	Likes       int64  `json:"likes"`
	LikedByUser bool   `json:"liked_by_user"`
}

type LoadParams struct {
	UserID uuid.UUID
	Order  string
	pages.PaginationParams
}

type PostRepo interface {
	Create(p Post) error
	Load(params LoadParams) ([]FeedPost, error)
	LoadByID(id uuid.UUID) (Post, error)
	Update(p Post) error
	Delete(p Post) error
}
