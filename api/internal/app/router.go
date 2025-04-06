package app

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/doktorupnos/crow/api/internal/respond"
)

type State struct {
	DB        *database.Queries
	Secret    string
	ExpiresIn time.Duration
}

func Router(s *State) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", Healthz)

	mux.HandleFunc("POST /api/login", WithError(s.BasicAuth(s.Login)))
	mux.HandleFunc("POST /api/logout", WithError(s.BasicAuth(s.Logout)))
	mux.HandleFunc("POST /api/admin/jwt", WithError(s.JWT(s.ValidateJWT)))

	userServer := &UserServer{
		Service:      &UserServicePG{db: s.DB},
		JWTSecret:    s.Secret,
		JWTExpiresIn: s.ExpiresIn,
	}
	mux.HandleFunc("POST /api/users", WithError(userServer.CreateUser))

	mux.HandleFunc("POST /api/posts", WithError(s.JWT(s.CreatePost)))
	mux.HandleFunc("GET /api/posts", WithError(s.JWT(s.GetPosts)))
	mux.HandleFunc("DELETE /api/posts/{id}", WithError(s.JWT(s.DeletePost)))

	mux.HandleFunc("POST /api/post_likes", WithError(s.JWT(s.CreateLike)))
	mux.HandleFunc("DELETE /api/post_likes", WithError(s.JWT(s.DeleteLike)))

	mux.HandleFunc("POST /api/follow", WithError(s.JWT(s.CreateFollow)))
	mux.HandleFunc("DELETE /api/follow", WithError(s.JWT(s.DeleteFollow)))
	mux.HandleFunc("GET /api/followers", WithError(s.JWT(s.GetFollowers)))
	mux.HandleFunc("GET /api/following", WithError(s.JWT(s.GetFollowing)))
	mux.HandleFunc("GET /api/followers_count", WithError(s.JWT(s.GetFollowerCount)))
	mux.HandleFunc("GET /api/following_count", WithError(s.JWT(s.GetFollowingCount)))

	mux.HandleFunc("GET /api/profile", WithError(s.JWT(s.Profile)))

	return Logger(mux)
}

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

type SpyResponseWriter struct {
	http.ResponseWriter
	Code int
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.Code = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

func (s *SpyResponseWriter) Header() http.Header {
	return s.ResponseWriter.Header()
}

func (s *SpyResponseWriter) Write(p []byte) (n int, err error) {
	return s.ResponseWriter.Write(p)
}

func Logger(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sw := &SpyResponseWriter{ResponseWriter: w}
		method := r.Method
		path := r.URL.Path
		handler.ServeHTTP(sw, r)
		log.Printf("%s %q %d", method, path, sw.Code)
	}
}
