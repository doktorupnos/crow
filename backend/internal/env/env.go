package env

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Env groups all the environment variables the server depends on
type Env struct {
	ServerAddr string
	CorsOrigin string
	DSN        string
}

// Load is a Env constructor.
// If a -local flag was specified when running the program then Load will depend on a .env file using godotenv.
func Load() (*Env, error) {
	local := flag.Bool("local", false, "Depend on a .env file for local development")
	flag.Parse()
	if *local {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

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

	return &Env{
		ServerAddr: serverAddr,
		CorsOrigin: corsOrigin,
		DSN:        dsn,
	}, nil
}

func envNotSet(name string) error {
	return fmt.Errorf("%s environment variable is not set", name)
}
