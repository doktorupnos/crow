package message

import (
	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/google/uuid"
)

type Repo interface {
	Create(*Message) error
	Load(LoadParams) ([]Message, error)
}

type LoadParams struct {
	Order     string
	ChannelID uuid.UUID
	pages.PaginationParams
}
