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

func (s *ChatServer) accept(conn *websocket.Conn) {
	log.Println("client connected:", conn.RemoteAddr())

	s.mu.Lock()
	s.conns[conn] = struct{}{}
	s.mu.Unlock()

	s.echo(conn)
}

func (s *ChatServer) echo(conn *websocket.Conn) {
	for {
		var message string
		err := websocket.Message.Receive(conn, &message)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("reading message: client disconnected: EOF")

				// TODO: Architect the behavior of closing connections for chat rooms
				// based on what the front-end needs.
				s.mu.Lock()
				delete(s.conns, conn)
				s.mu.Unlock()
				break
			}

			log.Println("reading message:", err)
			continue
		}

		log.Println("from", conn.RemoteAddr(), ":", message)
		websocket.Message.Send(conn, fmt.Sprintf("Echo: %s", message))
	}
}
