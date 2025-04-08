package app

import (
	"errors"
	"log"
	"net/http"

	"github.com/doktorupnos/crow/api/internal/respond"
)

type ErrorHandler func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Code int
	Err  error
}

func (e APIError) Error() string {
	return string(e.Err.Error())
}

func WithError(handler ErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		v, ok := err.(APIError)
		if !ok {
			const code = http.StatusInternalServerError
			respond.Error(w, code, errors.New(http.StatusText(code)))
			return
		}

		respond.Error(w, v.Code, v.Err)
	}
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func CORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func Logger(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sw := &StatusCodeResponseWriter{ResponseWriter: w}
		method := r.Method
		path := r.URL.Path
		handler.ServeHTTP(sw, r)
		log.Printf("%s %q %d", method, path, sw.Code)
	}
}

type StatusCodeResponseWriter struct {
	http.ResponseWriter
	Code int
}

func (s *StatusCodeResponseWriter) WriteHeader(statusCode int) {
	s.Code = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

func (s *StatusCodeResponseWriter) Header() http.Header {
	return s.ResponseWriter.Header()
}

func (s *StatusCodeResponseWriter) Write(p []byte) (n int, err error) {
	return s.ResponseWriter.Write(p)
}
