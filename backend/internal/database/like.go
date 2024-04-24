package database

import (
	"github.com/doktorupnos/crow/backend/internal/like"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormLikeRepo struct {
	db *gorm.DB
}

func NewGormLikeRepo(db *gorm.DB) *GormLikeRepo {
	return &GormLikeRepo{db}
}

func (r *GormLikeRepo) Create(l like.Like) error {
	return r.db.Create(&l).Error
}

func (r *GormLikeRepo) Load() ([]like.Like, error) {
	var likes []like.Like
	err := r.db.Find(&likes).Error
	return likes, err
}

func (r *GormLikeRepo) Single(userID, postID uuid.UUID) (like.Like, error) {
	var l like.Like
	err := r.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&l).Error
	return l, err
}

func (r *GormLikeRepo) Delete(l like.Like) error {
	return r.db.Delete(&l).Error
}
