package post

import (
	"github.com/doktorupnos/crow/backend/internal/model"
	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	Body  string      `json:"body" gorm:"not null"`
	Likes []user.User `json:"-" gorm:"many2many:post_likes"`
	User  user.User   `json:"-" gorm:"foreignKey:UserID; not null;constraint:onDelete:CASCADE"`
	model.Model
	UserID uuid.UUID `json:"user_id"`
}

func (p Post) BeforeDelete(db *gorm.DB) error {
	q := `DELETE FROM post_likes WHERE post_id = ?`
	return db.Exec(q, p.ID).Error
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
