package database

import "github.com/google/uuid"

type Post struct {
	Model
	Body string `json:"body" gorm:"size:255;not null"`

	UserID uuid.UUID
	User   User `json:"-" gorm:"foreignKey:UserID;not null;constraint:OnDelete:CASCADE"`
}
