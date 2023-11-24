package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdatePost(db *gorm.DB, postID, userID uuid.UUID, body string) (Post, error) {
	var post Post
	if err := db.Where("id = ? AND user_id = ?", postID, userID).First(&post).Error; err != nil {
		return Post{}, err
	}

	post.Body = body
	if err := db.Save(&post).Error; err != nil {
		return Post{}, err
	}

	return post, nil
}
