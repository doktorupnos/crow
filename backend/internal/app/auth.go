package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/channel"
	"github.com/doktorupnos/crow/backend/internal/jwt"
	"github.com/doktorupnos/crow/backend/internal/passwd"
	"github.com/doktorupnos/crow/backend/internal/respond"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}

const (
	ErrMissingAuthBasicHeader = AuthError("missing Authorization Basic header")
	ErrWrongPassword          = AuthError("wrong password")
)

type authedHandler func(w http.ResponseWriter, r *http.Request, u user.User)

func (app *App) BasicAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			respond.Error(w, http.StatusUnauthorized, ErrMissingAuthBasicHeader)
			return
		}

		u, err := app.userService.GetByName(username)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		if !passwd.Match(u.Password, password) {
			respond.Error(w, http.StatusUnauthorized, ErrWrongPassword)
			return
		}

		handler(w, r, u)
	}
}

func (app *App) JWT(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := app.extractUser(w, r)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
		}
		handler(w, r, u)
	}
}

func (app *App) extractUser(_ http.ResponseWriter, r *http.Request) (user.User, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return user.User{}, err
	}

	tokenString := c.Value

	token, err := jwt.Parse(app.Env.JWT.Secret, tokenString)
	if err != nil {
		return user.User{}, err
	}

	userIDString, err := token.Claims.GetSubject()
	if err != nil {
		return user.User{}, err
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return user.User{}, err
	}

	u, err := app.userService.GetByID(userID)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func (app *App) ValidateJWT(w http.ResponseWriter, r *http.Request, u user.User) {
	w.WriteHeader(http.StatusOK)
}

func (app *App) WS(f func(conn *websocket.Conn, u user.User, c channel.Channel)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := app.extractUser(w, r)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		channelName := r.URL.Query().Get("channel")
		if channelName == "" {
			respond.Error(w, http.StatusBadRequest, fmt.Errorf("channel: %q not found", channelName))
			return
		}

		c, err := app.channelService.GetByName(channelName)
		if err != nil {
			respond.Error(w, http.StatusBadRequest, err)
			return
		}

		conn, err := app.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			respond.Error(w, http.StatusInternalServerError, err)
			return
		}
		defer conn.Close()

		f(conn, u, c)
	}
}
