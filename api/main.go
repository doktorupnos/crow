package main

import (
	"github.com/doktorupnos/crow/api/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	app.Run()
}
