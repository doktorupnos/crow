package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/doktorupnos/crow/api/internal/respond"
)

type Server struct {
	service   Service
	secret    string
	expiresIn time.Duration
}

func NewServer(service Service, secret string, expiresIn time.Duration) *Server {
	return &Server{
		service:   service,
		secret:    secret,
		expiresIn: expiresIn,
	}
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) error {
	var req CreateRequest
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		return err
	}

	user, err := s.service.Create(r.Context(), req)
	if err != nil {
		return err
	}

	respond.JWT(w, http.StatusCreated, s.secret, user.ID.String(), s.expiresIn)
	return nil
}
