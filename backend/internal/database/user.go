package database

import (
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
