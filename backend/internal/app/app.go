package app

import (
	"time"

	"gorm.io/gorm"
)

// App is used to implement stateful http handlers.
type App struct {
	// DB is an open database connection.
	DB *gorm.DB
	// JWT_SECRET is the secret used to sign tokens.
	JWT_SECRET string
	// JWT_EXPIRES_IN_MINUTES is the duration a token will expire in.
	JWT_EXPIRES_IN_MINUTES time.Duration
	// ORIGIN is the URI the CORS headers will allow
	ORIGIN string
}
