package follow

import (
	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type Repo interface {
	Follow(u, o user.User) error
	Unfollow(u, o user.User) error
	Following(p LoadParams) ([]Follow, error)
	Followers(p LoadParams) ([]Follow, error)
	FollowingCount(u user.User) (int, error)
	FollowersCount(u user.User) (int, error)
	FollowsUser(u, t user.User) (bool, error)
}

type LoadParams struct {
	UserID uuid.UUID
	pages.PaginationParams
}
