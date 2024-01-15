// Package database encapsulates the Database Layer of GORM (GO's most popular ORM)
package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/doktorupnos/crow/backend/internal/like"
	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/doktorupnos/crow/backend/internal/user"
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

	if err := db.SetupJoinTable(&post.Post{}, "Likes", &like.PostLike{}); err != nil {
		log.Println("Failed to setup the Likes join table")
		return err
	}

	return nil
}
