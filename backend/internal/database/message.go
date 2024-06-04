package database

import (
	"github.com/doktorupnos/crow/backend/internal/message"
	"gorm.io/gorm"
)

type GormMessageRepo struct {
	db *gorm.DB
}

var _ message.Repo = &GormMessageRepo{}

func NewGormMessageRepo(db *gorm.DB) *GormMessageRepo {
	return &GormMessageRepo{db}
}

func (r *GormMessageRepo) Create(m *message.Message) error {
	return r.db.Create(m).Error
}
