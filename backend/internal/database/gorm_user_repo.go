package database

import (
	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) *GormUserRepo {
	return &GormUserRepo{db}
}

func (r *GormUserRepo) Create(u user.User) (uuid.UUID, error) {
	err := r.db.Create(&u).Error
	return u.ID, err
}

func (r *GormUserRepo) GetAll() ([]user.User, error) {
	var users []user.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepo) GetByName(name string) (user.User, error) {
	var u user.User
	if err := r.db.Where("name = ?", name).First(&u).Error; err != nil {
		return user.User{}, err
	}
	return u, nil
}

func (r *GormUserRepo) GetByID(id uuid.UUID) (user.User, error) {
	var u user.User
	if err := r.db.First(&u, id).Error; err != nil {
		return user.User{}, err
	}
	return u, nil
}

func (r *GormUserRepo) Update(u user.User) error {
	return r.db.Save(&u).Error
}

func (r *GormUserRepo) Delete(u user.User) error {
	return r.db.Delete(&u).Error
}

func (r *GormUserRepo) Follow(u, o user.User) error {
	return r.db.Model(&u).Association("Follows").Append(&o)
}

func (r *GormUserRepo) Unfollow(u, o user.User) error {
	return r.db.Model(&u).Association("Follows").Delete(&o)
}

func (r *GormUserRepo) Following(p user.LoadParams) ([]user.Follow, error) {
	q := `SELECT uf.follow_id, u.name
FROM user_follows uf JOIN users u ON uf.follow_id = u.id
WHERE uf.user_id = ?`

	rows, err := r.db.Scopes(pages.Paginate(p.PaginationParams)).Raw(q, p.UserID).Rows()
	if err != nil {
		return nil, err
	}

	following := []user.Follow{}
	follow := user.Follow{}
	for rows.Next() {
		err := rows.Scan(&follow.ID, &follow.Name)
		if err != nil {
			return nil, err
		}
		following = append(following, follow)
	}

	return following, nil
}

func (r *GormUserRepo) Followers(p user.LoadParams) ([]user.Follow, error) {
	q := `SELECT uf.user_id, u.name
FROM user_follows uf JOIN users u ON uf.user_id = u.id
WHERE uf.follow_id = ?`

	rows, err := r.db.Scopes(pages.Paginate(p.PaginationParams)).Raw(q, p.UserID).Rows()
	if err != nil {
		return nil, err
	}

	followers := []user.Follow{}
	follow := user.Follow{}
	for rows.Next() {
		err := rows.Scan(&follow.ID, &follow.Name)
		if err != nil {
			return nil, err
		}
		followers = append(followers, follow)
	}

	return followers, nil
}

func (r *GormUserRepo) FollowingCount(u user.User) (int, error) {
	q := `SELECT COUNT(*)
FROM user_follows
WHERE user_id = ?`

	var following int
	err := r.db.Raw(q, u.ID).Scan(&following).Error
	if err != nil {
		return 0, err
	}
	return following, nil
}

func (r *GormUserRepo) FollowersCount(u user.User) (int, error) {
	q := `SELECT COUNT(*)
FROM user_follows
WHERE follow_id = ?`

	var followers int
	err := r.db.Raw(q, u.ID).Scan(&followers).Error
	if err != nil {
		return 0, err
	}
	return followers, nil
}

func (r *GormUserRepo) FollowsUser(u, t user.User) (bool, error) {
	q := `SELECT COUNT(*)
  FROM user_follows
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
