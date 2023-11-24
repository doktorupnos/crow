package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, userID uuid.UUID) (User, error) {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func GetUserByName(db *gorm.DB, name string) (User, error) {
	var user User
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
