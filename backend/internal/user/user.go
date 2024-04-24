package user

import (
	"github.com/doktorupnos/crow/backend/internal/model"
)

type User struct {
	Name     string  `json:"name" gorm:"unique;not null"`
	Password string  `json:"-"    gorm:"not null"`
	Follows  []*User `gorm:"many2many:follows"`
	model.Model
}
