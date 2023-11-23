// Package database encapsulates the ORM
package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect creates a new Gorm database connection from dsn (Data Source Name).
// Connect pings the database to ensure a connection and performs an auto migration.
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = ping(db)
	if err != nil {
		return nil, err
	}

	err = automigrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ping(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	return err
}

func automigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("failed to migrate User type: %s", err)
	}

	return nil
}
