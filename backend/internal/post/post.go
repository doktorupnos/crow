package post

import (
	"github.com/doktorupnos/crow/backend/internal/model"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	model.Model

	Body   string    `json:"body"    gorm:"not null"`
	UserID uuid.UUID `json:"user_id"`
	User   user.User `json:"-"       gorm:"foreignKey:UserID; not null;constraint:onDelete:CASCADE"`
}

type FeedPost struct {
	Post
	UserName string `json:"user_name"`
}

type LoadParams struct {
	Order string
	PaginationParams
}

type PostRepo interface {
	Create(p Post) error
	Load(params LoadParams) ([]FeedPost, error)
	LoadByID(id uuid.UUID) (Post, error)
	Update(p Post) error
	Delete(p Post) error
}

type PaginationParams struct {
	PageNumber int
	PageSize   int
}

func Paginate(params PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (params.PageNumber - 1) * params.PageSize
		return db.Offset(offset).Limit(params.PageSize)
	}
}
