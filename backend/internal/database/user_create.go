package database

import (
	"gorm.io/gorm"
)

type CreateUserParams struct {
	Name     string
	Password string
}

func CreateUser(db *gorm.DB, params CreateUserParams) (User, error) {
	user := User{
		Name:     params.Name,
		Password: params.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}
