package main

import (
	"flag"
	"log"

	"github.com/doktorupnos/crow/backend/internal/app"
	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/doktorupnos/crow/backend/internal/env"
	"github.com/joho/godotenv"
)

func main() {
	local := flag.Bool("local", false, "Depend on a .env file for local development")
	flag.Parse()
	if *local {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	env, err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(env.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(env.Server.Addr)
	app.New(env, db).Run()
}
