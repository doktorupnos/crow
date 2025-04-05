package app

import (
	"log"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/api/internal/database"
)

type State struct {
	DB        database.Querier
	Secret    string
	ExpiresIn time.Duration
}

func Router(s *State) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", Healthz)

	mux.HandleFunc("POST /api/login", s.BasicAuth(s.Login))
	mux.HandleFunc("POST /api/logout", s.BasicAuth(s.Logout))
	mux.HandleFunc("POST /api/admin/jwt", s.JWT(s.ValidateJWT))

	mux.HandleFunc("POST /api/users", s.CreateUser)

	mux.HandleFunc("POST /api/posts", s.JWT(s.CreatePost))
	mux.HandleFunc("GET /api/posts", s.JWT(s.GetPosts))
	mux.HandleFunc("DELETE /api/posts/{id}", s.JWT(s.DeletePost))

	mux.HandleFunc("POST /api/post_likes", s.JWT(s.CreateLike))
	mux.HandleFunc("DELETE /api/post_likes", s.JWT(s.DeleteLike))

	mux.HandleFunc("POST /api/follow", s.JWT(s.CreateFollow))
	mux.HandleFunc("DELETE /api/follow", s.JWT(s.DeleteFollow))
	mux.HandleFunc("GET /api/followers", s.JWT(s.GetFollowers))
	mux.HandleFunc("GET /api/following", s.JWT(s.GetFollowing))
	mux.HandleFunc("GET /api/followers_count", s.JWT(s.GetFollowerCount))
	mux.HandleFunc("GET /api/following_count", s.JWT(s.GetFollowingCount))

	mux.HandleFunc("GET /api/profile", s.JWT(s.Profile))

	return Logger(mux)
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
