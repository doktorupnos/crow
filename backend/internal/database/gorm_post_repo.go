package database

import (
	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/google/uuid"
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

func (r *GormPostRepo) GetByID(id uuid.UUID) (post.Post, error) {
	var p post.Post
	err := r.db.First(&p, id).Error
	return p, err
}

func (r *GormPostRepo) Update(p post.Post) error {
	return r.db.Save(&p).Error
}

func (r *GormPostRepo) Delete(p post.Post) error {
	return r.db.Delete(&p).Error
}
