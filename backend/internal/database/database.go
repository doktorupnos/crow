// Package database encapsulates the Database Layer of GORM (GO's most popular ORM)
package database

import (
	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/doktorupnos/crow/backend/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect creates a new connection with a Postgres database through GORM.
func Connect(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := automigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func automigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&post.Post{}); err != nil {
		return err
	}

	return nil
}
