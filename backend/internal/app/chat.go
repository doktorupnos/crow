package app

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/doktorupnos/crow/backend/internal/channel"
	"github.com/doktorupnos/crow/backend/internal/message"
	"github.com/doktorupnos/crow/backend/internal/respond"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/gorilla/websocket"
)

type ChatServer struct {
	conns    map[*websocket.Conn]struct{}
	ms       *message.Service
	upgrader websocket.Upgrader
	mu       sync.Mutex
}

func NewChatServer(ms *message.Service) *ChatServer {
	return &ChatServer{
		conns: make(map[*websocket.Conn]struct{}),
		ms:    ms,
	}
}

func (s *ChatServer) upgrade(f func(*websocket.Conn)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			respond.Error(w, http.StatusInternalServerError, err)
			return
		}
		defer conn.Close()
		f(conn)
	}
}

// accept persists the incoming websocket connection to the map of known connections
// which is used for broadcasting.
func (s *ChatServer) accept(conn *websocket.Conn) {
	log.Println("client connected:", conn.RemoteAddr())

	s.mu.Lock()
	defer s.mu.Unlock()
	s.conns[conn] = struct{}{}
}

// disconnect deletes a disconnected websocket connection from the map of known connections
// Disconnecting is typically associated with an io.EOF read error but can also be used for connection management
// when repeated broadcast errors occur.
func (s *ChatServer) disconnect(conn *websocket.Conn) {
	log.Println("client", conn.RemoteAddr(), "disconnected")

	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.conns, conn)
}

// echo is an example handler documenting the websocket send-receive model.
// It should be used by the front-end as a way to test single message communication
func (s *ChatServer) echo(conn *websocket.Conn) {
	s.accept(conn)

	for {
		_, body, err := conn.ReadMessage()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// TODO: Architect the behavior of closing connections for chat rooms
				// based on what the front-end needs.

				s.disconnect(conn)
				break
			}

			log.Println("reading message:", err)
			continue
		}

		body = []byte(fmt.Sprintf("Echo: %s", body))
		conn.WriteMessage(websocket.TextMessage, body)
	}
}

// cosmos is an example handler documenting websocket broadcasting
func (s *ChatServer) cosmos(conn *websocket.Conn, u user.User, c channel.Channel) {
	s.accept(conn)

	for {
		_, body, err := conn.ReadMessage()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("reading message: client disconnected: EOF")

				s.disconnect(conn)
				break
			}

			log.Println("reading message:", err)
			continue
		}

		log.Println("from", conn.RemoteAddr(), "by", u.Name, ":", string(body))
		s.ms.Create(message.CreateParams{
			Body:    string(body),
			User:    u,
			Channel: c,
		})
		s.broadcast(body)
	}
}

// feed is an example handler documenting a websocket subscription feed.
// A client "subscribes" to the feed and receives real-time updates.
func (s *ChatServer) feed(conn *websocket.Conn) {
	s.accept(conn)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; true; <-ticker.C {
		message := fmt.Sprintf("%d\n", time.Now().UnixNano())
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}

func (s *ChatServer) broadcast(body []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for conn := range s.conns {
		go func(conn *websocket.Conn) {
			if err := conn.WriteMessage(websocket.TextMessage, body); err != nil {
				log.Println("broadcasting to", conn.RemoteAddr(), ":", err)
			}
		}(conn)
	}
}
