package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreatePostParams struct {
	Body   string
	UserID uuid.UUID
}

func CreatePost(db *gorm.DB, params CreatePostParams) (Post, error) {
	post := Post{
		Body:   params.Body,
		UserID: params.UserID,
	}

	if err := db.Create(&post).Error; err != nil {
		return Post{}, err
	}

	return post, nil
}
