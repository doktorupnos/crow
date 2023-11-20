package main

import (
	"time"

	"github.com/google/uuid"
)

// Model contains common columns shared by all database tables.
type Model struct {
	// ID is a global version 4 unique identifier serving as the primary key
	ID uuid.UUID `json:"id"         gorm:"type:uuid; primaryKey;default: gen_random_uuid()"`
	// CreatedAt is the time the record was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the time the record was last updated
	UpdatedAt time.Time `json:"updated_at"`
}
