package post

import (
	"os/user"

	"github.com/doktorupnos/crow/backend/internal/model"
	"github.com/google/uuid"
)

type Post struct {
	model.Model
	Body string `json:"body" gorm:"not null"`

	UserID uuid.UUID `json:"user_id"`
	User   user.User `json:"-"       gorm:"foreignKey:UserID; not null;constraint:onDelete:CASCADE"`
}

type PostRepo interface {
	Create(p Post) error
}
