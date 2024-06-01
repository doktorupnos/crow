package app

import (
	"errors"
	"fmt"
	"io"
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

type ChatServer struct {
	conns map[*websocket.Conn]struct{}
	mu    sync.Mutex
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		conns: make(map[*websocket.Conn]struct{}),
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

// echo is an example handler documenting the websocket send-receive model
func (s *ChatServer) echo(conn *websocket.Conn) {
	s.accept(conn)

	for {
		var message string
		err := websocket.Message.Receive(conn, &message)
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

		log.Println("from", conn.RemoteAddr(), ":", message)
		websocket.Message.Send(conn, fmt.Sprintf("Echo: %s", message))
	}
}

// world is an example handler documenting websocket broadcasting
func (s *ChatServer) world(conn *websocket.Conn) {
	s.accept(conn)

	for {
		var message string
		err := websocket.Message.Receive(conn, &message)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("reading message: client disconnected: EOF")

				s.mu.Lock()
				defer s.mu.Unlock()
				delete(s.conns, conn)

				break
			}

			log.Println("reading message:", err)
			continue
		}

		log.Println("from", conn.RemoteAddr(), ":", message)
		s.broadcast(message)
	}
}

func (s *ChatServer) broadcast(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for conn := range s.conns {
		go func(conn *websocket.Conn) {
			if err := websocket.Message.Send(conn, message); err != nil {
				log.Println("broadcasting to", conn.RemoteAddr(), ":", err)
			}
		}(conn)
	}
}
