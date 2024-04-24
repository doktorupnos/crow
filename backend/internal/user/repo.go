package user

import (
	"github.com/doktorupnos/crow/backend/internal/follow"
	"github.com/google/uuid"
)

type Repo interface {
	Create(u User) (uuid.UUID, error)
	GetAll() ([]User, error)
	GetByName(name string) (User, error)
	GetByID(id uuid.UUID) (User, error)
	Update(u User) error
	Delete(u User) error

	Follow(u, o User) error
	Unfollow(u, o User) error

	Following(p LoadParams) ([]follow.Follow, error)
	Followers(p LoadParams) ([]follow.Follow, error)
	FollowingCount(u User) (int, error)
	FollowersCount(u User) (int, error)
	FollowsUser(u, t User) (bool, error)
}
