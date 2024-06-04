package database

import (
	"github.com/doktorupnos/crow/backend/internal/channel"
	"gorm.io/gorm"
)

type GormChannelRepo struct {
	db *gorm.DB
}

var _ channel.Repo = &GormChannelRepo{}

func NewGormChannelRepo(db *gorm.DB) *GormChannelRepo {
	return &GormChannelRepo{db}
}

func (r *GormChannelRepo) GetByName(name string) (channel.Channel, error) {
	var c channel.Channel
	if err := r.db.Where("name = ?", name).First(&c).Error; err != nil {
		return channel.Channel{}, err
	}
	return c, nil
}
