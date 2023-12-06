package like

import "github.com/google/uuid"

type PostLike struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	PostID    uuid.UUID `gorm:"primaryKey"`
	CreatedAt int64     `gorm:"autoCreateTime"`
}
