package database

type User struct {
	Model
	Name     string `json:"name" gorm:"size:20;unique;not null"`
	Password string `json:"-"    gorm:"size:72;not null"`
}
