package app

import (
	"net/http"
	"time"

	"github.com/doktorupnos/crow/api/internal/app/user"
	"github.com/doktorupnos/crow/api/internal/database"
)

type State struct {
	db        *database.Queries
	secret    string
	expiresIn time.Duration
}

func Router(state *State) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", Healthz)

	mux.HandleFunc("POST /api/login", WithError(state.BasicAuth(state.Login)))
	mux.HandleFunc("POST /api/logout", WithError(state.BasicAuth(state.Logout)))
	mux.HandleFunc("POST /api/admin/jwt", WithError(state.JWT(state.ValidateJWT)))

	userServer := user.NewServer(user.NewPostgresService(state.db), state.secret, state.expiresIn)
	mux.HandleFunc("POST /api/users", WithError(userServer.Create))

	mux.HandleFunc("POST /api/posts", WithError(state.JWT(state.CreatePost)))
	mux.HandleFunc("GET /api/posts", WithError(state.JWT(state.GetPosts)))
	mux.HandleFunc("DELETE /api/posts/{id}", WithError(state.JWT(state.DeletePost)))

	mux.HandleFunc("POST /api/post_likes", WithError(state.JWT(state.CreateLike)))
	mux.HandleFunc("DELETE /api/post_likes", WithError(state.JWT(state.DeleteLike)))

	mux.HandleFunc("POST /api/follow", WithError(state.JWT(state.CreateFollow)))
	mux.HandleFunc("DELETE /api/follow", WithError(state.JWT(state.DeleteFollow)))
	mux.HandleFunc("GET /api/followers", WithError(state.JWT(state.GetFollowers)))
	mux.HandleFunc("GET /api/following", WithError(state.JWT(state.GetFollowing)))
	mux.HandleFunc("GET /api/followers_count", WithError(state.JWT(state.GetFollowerCount)))
	mux.HandleFunc("GET /api/following_count", WithError(state.JWT(state.GetFollowingCount)))

	mux.HandleFunc("GET /api/profile", WithError(state.JWT(state.Profile)))

	return Logger(CORS(mux))
}
