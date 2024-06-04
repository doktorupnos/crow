package message

import (
	"github.com/doktorupnos/crow/backend/internal/channel"
	"github.com/doktorupnos/crow/backend/internal/user"
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
