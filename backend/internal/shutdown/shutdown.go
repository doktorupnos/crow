package shutdown

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ListenAndServe(s *http.Server, timeout time.Duration) {
	errs := make(chan error, 1)
	go func(s *http.Server) {
		errs <- s.ListenAndServe()
	}(s)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var count int

	if err := s.Shutdown(ctx); err != nil {
		log.Println("http.Server.Shutdown:", err)
		count++
	}
	if err := s.Close(); err != nil {
		log.Println("http.Server.Close:", err)
		count++
	}

	if err := <-errs; !errors.Is(err, http.ErrServerClosed) {
		log.Println("http.Server.ListenAndServe:", err)
		count++
	}

	if count == 0 {
		log.Println("Graceful shutdown")
	} else {
		log.Println("Ungraceful shutdown")
	}
}
