package user

import (
	"github.com/doktorupnos/crow/backend/internal/model"
	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/google/uuid"
)

type User struct {
	Name     string  `json:"name" gorm:"unique;not null"`
	Password string  `json:"-"    gorm:"not null"`
	Follows  []*User `gorm:"many2many:user_follows"`
	model.Model
}

type Follow struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

type LoadParams struct {
	UserID uuid.UUID
	pages.PaginationParams
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

	Following(p LoadParams) ([]Follow, error)
	Followers(p LoadParams) ([]Follow, error)
	FollowingCount(u User) (int, error)
	FollowersCount(u User) (int, error)
	FollowsUser(u, t User) (bool, error)
}
