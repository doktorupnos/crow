package message

import (
	"net/http"

	"github.com/doktorupnos/crow/backend/internal/channel"
	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type Service struct {
	r Repo
}

func NewService(r Repo) *Service {
	return &Service{r}
}

type CreateParams struct {
	Body    string
	Channel channel.Channel
	User    user.User
}

func (s *Service) Create(p CreateParams) (*Message, error) {
	message := &Message{
		Body:      p.Body,
		User:      p.User,
		UserID:    p.User.ID,
		Channel:   p.Channel,
		ChannelID: p.Channel.ID,
	}

	err := s.r.Create(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (s *Service) Load(r *http.Request, defaultPageSize int, channelID uuid.UUID) ([]Message, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}
	order := "desc"

	return s.r.Load(LoadParams{
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
		Order:     order,
		ChannelID: channelID,
	})
}
