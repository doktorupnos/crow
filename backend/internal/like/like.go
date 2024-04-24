package like

import "github.com/google/uuid"

type Like struct {
	UserID uuid.UUID `gorm:"primaryKey"`
	PostID uuid.UUID `gorm:"primaryKey"`
	// CreatedAt int64     `gorm:"autoCreateTime"`
}
