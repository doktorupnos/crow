// Package database encapsulates the Database Layer of GORM (GO's most popular ORM)
package database

import (
	"log"

	"github.com/doktorupnos/crow/backend/internal/like"
	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/doktorupnos/crow/backend/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect creates a new connection with a Postgres database through GORM.
func Connect(dataSourceName string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), config)
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

	if err := db.SetupJoinTable(&post.Post{}, "Likes", &like.Like{}); err != nil {
		log.Println("Failed to setup the Likes join table")
		return err
	}

	return nil
}
