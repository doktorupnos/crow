// Package database encapsulates the Database Layer of GORM (GO's most popular ORM)
package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect creates a new connection with a Postgres database through GORM.
func Connect(dataSourceName string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
}
