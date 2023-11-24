// The environment variable names are not final and should be set from Kubernetes
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Env contains environment variables that must be set for the server to run.
type Env struct {
	// The port for the server to listen on
	PORT string
	// The Data Source Name of the database
	DSN string
	// The secret to sign JWTs
	JWT_SECRET string

	// The duraction a JWT access token expires in
	JWT_EXPIRES_IN_MINUTES time.Duration
}

func loadEnv() (*Env, error) {
	local := flag.Bool("local", true, "Depend on a .env file for local development")
	flag.Parse()

	if *local {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load .env : %s", err)
		}
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		return nil, errors.New("PORT environment variable is not set")
	}

	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		return nil, errors.New("DSN environment variable is not set")
	}

	jwtSecret, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		return nil, errors.New("JWT_SECRET environment variable is not set")
	}

	jwtExpiresInMinutesString, ok := os.LookupEnv("JWT_EXPIRES_IN_MINUTES")
	if !ok {
		return nil, errors.New("JWT_EXPIRES_IN_MINUTES enviornment variable is not set")
	}
	jwtExpiresInMinutes, err := strconv.Atoi(jwtExpiresInMinutesString)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse JWT_EXPIRES_IN_MINUTES environment variable : %s",
			err,
		)
	}

	return &Env{
		PORT:                   port,
		DSN:                    dsn,
		JWT_SECRET:             jwtSecret,
		JWT_EXPIRES_IN_MINUTES: time.Minute * time.Duration(jwtExpiresInMinutes),
	}, nil
}
