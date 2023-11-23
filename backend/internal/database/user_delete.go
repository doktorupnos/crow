package database

import "gorm.io/gorm"

func DeleteUser(db *gorm.DB, user User) error {
	return db.Delete(&user).Error
}
