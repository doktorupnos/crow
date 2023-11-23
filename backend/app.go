package main

import "gorm.io/gorm"

// App is used to implement stateful http handlers.
type App struct {
	DB         *gorm.DB
	JWT_SECRET string
}
