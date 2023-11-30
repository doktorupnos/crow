package env

import (
	"fmt"
	"os"
	"time"
)

// Env groups all the environment variables the server depends on
type Env struct {
	ServerAddr  string
	CorsOrigin  string
	DSN         string
	JwtSecret   string
	JwtLifetime time.Duration
}

func Load() (*Env, error) {
	serverAddr, ok := os.LookupEnv("ADDR")
	if !ok {
		return nil, envNotSet("ADDR")
	}

	corsOrigin, ok := os.LookupEnv("CORS_ORIGIN")
	if !ok {
		return nil, envNotSet("CORS_ORIGIN")
	}

	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		return nil, envNotSet("DSN")
	}

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, envNotSet("JWT_SECRET")
	}

	jwtLifetimeString, ok := os.LookupEnv("JWT_LIFETIME")
	if !ok {
		return nil, envNotSet("JWT_LIFETIME")
	}
	jwtLifetime, err := time.ParseDuration(jwtLifetimeString)
	if err != nil {
		return nil, err
	}

	return &Env{
		ServerAddr:  serverAddr,
		CorsOrigin:  corsOrigin,
		DSN:         dsn,
		JwtSecret:   jwtSecret,
		JwtLifetime: jwtLifetime,
	}, nil
}

func envNotSet(name string) error {
	return fmt.Errorf("%s environment variable is not set", name)
}
