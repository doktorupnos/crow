package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func GetAllPostsByUserID(db *gorm.DB, userID uuid.UUID) ([]Post, error) {
	var posts []Post
	if err := db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostByID(db *gorm.DB, postID uuid.UUID) (Post, error) {
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		return Post{}, err
	}
	return post, nil
}

func GetAllPostsByUserName(db *gorm.DB, name string) ([]Post, error) {
	var posts []Post
	err := db.Model(&Post{}).
		Select("posts.*").
		Joins("join users on posts.user_id = users.id").
		Where("users.name = ?", name).
		Scan(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
