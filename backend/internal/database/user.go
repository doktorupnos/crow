package database

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Model
	Name           string `json:"name" gorm:"size:20;unique;not null"`
	HashedPassword string `json:"-"    gorm:"size:64;not null"`
}

type CreateUserParams struct {
	Name           string
	HashedPassword string
}

func CreateUser(db *gorm.DB, params CreateUserParams) (User, error) {
	user := User{
		Name:           params.Name,
		HashedPassword: params.HashedPassword,
	}

	if err := db.Create(&user).Error; err != nil {
		return User{}, fmt.Errorf("failed to cretae user : %s", err)
	}

	return user, nil
}

func GetUserByName(db *gorm.DB, name string) (User, error) {
	var user User
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		return User{}, fmt.Errorf("failed to retrieve user with name %s : %s", name, err)
	}

	return user, nil
}

func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve users : %s", err)
	}

	return users, nil
}
