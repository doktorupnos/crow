package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Full Rewrite!")

	mux := http.NewServeMux()
	mux.Handle("GET /healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	const addr = ":8000"
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Println("api serving on", addr)
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
