package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UpdateUserParams struct {
	Name     string
	Password string
}

func UpdateUser(db *gorm.DB, userID uuid.UUID, params UpdateUserParams) (User, error) {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return User{}, err
	}

	user.Name = params.Name
	user.Password = params.Password

	if err := db.Save(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
