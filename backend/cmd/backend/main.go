package main

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/doktorupnos/crow/backend/internal/app"
	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/doktorupnos/crow/backend/internal/env"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
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

	data, err := json.MarshalIndent(env, "", "\t")
	if err != nil {
		log.Println("Marshalling env:", env)
	} else {
		log.Println(`"Env":`, string(data))
	}

	db, err := database.Connect(env.Database.DSN, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Serving from:", env.Server.Addr)
	app.New(env, db).Run()
}
