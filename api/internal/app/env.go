package app

import (
	"fmt"
	"os"
	"time"
)

const (
	DSN      = "DSN"
	Secret   = "JWT_SECRET"
	Lifetime = "JWT_LIFETIME"
)

type Env struct {
	DSN string
	JWT JWT
}

func NewEnv() (*Env, error) {
	dsn, ok := os.LookupEnv(DSN)
	if !ok {
		return notSet(DSN)
	}

	secret, ok := os.LookupEnv(Secret)
	if !ok {
		return notSet(Secret)
	}

	expiresStr, ok := os.LookupEnv(Lifetime)
	if !ok {
		return notSet(Lifetime)
	}
	expiresIn, err := time.ParseDuration(expiresStr)
	if err != nil {
		return nil, err
	}

	return &Env{
		DSN: dsn,
		JWT: JWT{
			Secret:    secret,
			ExpiresIn: expiresIn,
		},
	}, nil
}

func notSet(key string) (*Env, error) {
	return nil, fmt.Errorf("environment variable %q not set", key)
}

type JWT struct {
	Secret    string
	ExpiresIn time.Duration
}
