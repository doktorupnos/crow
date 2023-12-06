package like

import "github.com/google/uuid"

type PostLike struct {
	UserID uuid.UUID `gorm:"primaryKey"`
	PostID uuid.UUID `gorm:"primaryKey"`
	// CreatedAt int64     `gorm:"autoCreateTime"`
}

type PostLikeRepo interface {
	Create(l PostLike) error
	Load() ([]PostLike, error)
	Single(userID, postID uuid.UUID) (PostLike, error)
	Delete(l PostLike) error
}
