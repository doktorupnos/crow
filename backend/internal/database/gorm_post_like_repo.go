package database

import (
	"github.com/doktorupnos/crow/backend/internal/like"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormPostLikeRepo struct {
	db *gorm.DB
}

func NewGormPostLikeRepo(db *gorm.DB) *GormPostLikeRepo {
	return &GormPostLikeRepo{db}
}

func (r *GormPostLikeRepo) Create(l like.PostLike) error {
	return r.db.Create(&l).Error
}

func (r *GormPostLikeRepo) Load() ([]like.PostLike, error) {
	var likes []like.PostLike
	err := r.db.Find(&likes).Error
	return likes, err
}

func (r *GormPostLikeRepo) Single(userID, postID uuid.UUID) (like.PostLike, error) {
	var l like.PostLike
	err := r.db.Where("user_id = ? AND post_id = ?", userID, postID).First(&l).Error
	return l, err
}

func (r *GormPostLikeRepo) Delete(l like.PostLike) error {
	return r.db.Delete(&l).Error
}
