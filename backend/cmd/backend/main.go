package main

import (
	"log"

	"github.com/doktorupnos/crow/backend/internal/app"
	"github.com/doktorupnos/crow/backend/internal/database"
	"github.com/doktorupnos/crow/backend/internal/env"
)

func main() {
	env, err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(env.DSN)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(env.ServerAddr)
	app.New(env, db).Run()
}
