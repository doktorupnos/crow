package model

import (
	"github.com/google/uuid"
)

// Model contains common columns shared by almost all database tables.
type Model struct {
	// ID is a global version 4 unique identifier serving as the primary key
	ID uuid.UUID `json:"id"         gorm:"type:uuid; primaryKey; default: gen_random_uuid()"`
	// CreatedAt is the time the record was created
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	// UpdatedAt is the time the record was last updated
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}
