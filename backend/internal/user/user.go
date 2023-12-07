package user

import (
	"github.com/doktorupnos/crow/backend/internal/model"
	"github.com/google/uuid"
)

type User struct {
	model.Model
	Name     string `json:"name" gorm:"unique;not null"`
	Password string `json:"-"    gorm:"not null"`

	Follows []*User `gorm:"many2many:user_follows"`
}

type UserRepo interface {
	Create(u User) (uuid.UUID, error)
	GetAll() ([]User, error)
	GetByName(name string) (User, error)
	GetByID(id uuid.UUID) (User, error)
	Update(u User) error
	Delete(u User) error

	Follow(u, o User) error
	Unfollow(u, o User) error
}
