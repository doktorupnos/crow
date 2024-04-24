package database

import (
	"github.com/doktorupnos/crow/backend/internal/follow"
	"github.com/doktorupnos/crow/backend/internal/user"
	"gorm.io/gorm"
)

type GormFollowRepo struct {
	db *gorm.DB
}

func NewGormFollowRepo(db *gorm.DB) *GormFollowRepo {
	return &GormFollowRepo{db}
}

func (r *GormFollowRepo) Follow(u, o user.User) error {
	return r.db.Model(&u).Association("Follows").Append(&o)
}

func (r *GormFollowRepo) Unfollow(u, o user.User) error {
	return r.db.Model(&u).Association("Follows").Delete(&o)
}

func (r *GormFollowRepo) Following(p follow.LoadParams) ([]follow.Follow, error) {
	q := `SELECT f.follow_id, u.name
  FROM users u JOIN follows f ON u.id = f.follow_id
  WHERE f.user_id = ?
  LIMIT ? OFFSET ?`

	limit := p.PageSize
	offset := limit * p.PageNumber

	rows, err := r.db.Raw(q, p.UserID, limit, offset).Rows()
	if err != nil {
		return nil, err
	}
	following := []follow.Follow{}

	follow := follow.Follow{}
	for rows.Next() {
		err := rows.Scan(&follow.ID, &follow.Name)
		if err != nil {
			return nil, err
		}
		following = append(following, follow)
	}

	return following, nil
}

func (r *GormFollowRepo) Followers(p follow.LoadParams) ([]follow.Follow, error) {
	q := `SELECT f.user_id, u.name
				FROM follows f JOIN users u ON f.user_id = u.id
				WHERE f.follow_id = ?
		    LIMIT ? OFFSET ?`

	limit := p.PageSize
	offset := limit * p.PageNumber

	rows, err := r.db.Raw(q, p.UserID, limit, offset).Rows()
	if err != nil {
		return nil, err
	}

	followers := []follow.Follow{}
	follow := follow.Follow{}
	for rows.Next() {
		err := rows.Scan(&follow.ID, &follow.Name)
		if err != nil {
			return nil, err
		}
		followers = append(followers, follow)
	}

	return followers, nil
}

func (r *GormFollowRepo) FollowingCount(u user.User) (int, error) {
	q := `SELECT COUNT(*)
FROM follows
WHERE user_id = ?`

	var following int
	err := r.db.Raw(q, u.ID).Scan(&following).Error
	if err != nil {
		return 0, err
	}
	return following, nil
}

func (r *GormFollowRepo) FollowersCount(u user.User) (int, error) {
	q := `SELECT COUNT(*)
FROM follows
WHERE follow_id = ?`

	var followers int
	err := r.db.Raw(q, u.ID).Scan(&followers).Error
	if err != nil {
		return 0, err
	}
	return followers, nil
}

func (r *GormFollowRepo) FollowsUser(u, t user.User) (bool, error) {
	q := `SELECT COUNT(*)
  FROM follows
  WHERE user_id = ? AND follow_id = ?`

	var count int
	if err := r.db.Raw(q, u.ID, t.ID).Scan(&count).Error; err != nil {
		return false, err
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}
