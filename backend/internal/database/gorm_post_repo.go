package database

import (
	"github.com/doktorupnos/crow/backend/internal/post"
	"gorm.io/gorm"
)

type GormPostRepo struct {
	db *gorm.DB
}

func NewGormPostRepo(db *gorm.DB) *GormPostRepo {
	return &GormPostRepo{db}
}

func (r *GormPostRepo) Create(p post.Post) error {
	return r.db.Create(&p).Error
}

func (r *GormPostRepo) GetAll() ([]post.Post, error) {
	var posts []post.Post
	err := r.db.Find(&posts).Error
	return posts, err
}
