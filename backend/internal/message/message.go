package message

import (
	"github.com/doktorupnos/crow/backend/internal/channel"
	"github.com/doktorupnos/crow/backend/internal/model"
	"github.com/doktorupnos/crow/backend/internal/user"

	"github.com/google/uuid"
)

type Message struct {
	Body    string          `json:"body" gorm:"not null"`
	Channel channel.Channel `json:"-" gorm:"foreignKey:ChannelID; not null;constraint:onDelete:CASCADE"`
	User    user.User       `json:"-" gorm:"foreignKey:UserID; not null;constraint:onDelete:CASCADE"`
	model.Model
	UserID    uuid.UUID `json:"user_id"`
	ChannelID uuid.UUID `json:"channel_id"`
}
