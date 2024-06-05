package database

import (
	"github.com/doktorupnos/crow/backend/internal/message"
	"github.com/doktorupnos/crow/backend/internal/pages"
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

func (r *GormMessageRepo) Load(p message.LoadParams) ([]message.Message, error) {
	var messages []message.Message

	err := r.db.
		Order("created_at "+p.Order).
		Scopes(pages.Paginate(p.PaginationParams)).
		Find(&messages).
		Where("channel_id = ?", p.ChannelID).
		Error
	if err != nil {
		return nil, err
	}

	return messages, nil
}
