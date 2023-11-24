package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeletePost(db *gorm.DB, postID, userID uuid.UUID) error {
	var post Post
	err := db.Where("user_id = ?", userID).First(&post, postID).Error
	if err != nil {
		return err
	}

	return db.Delete(&post).Error
}
