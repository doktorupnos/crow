package channel

import "github.com/doktorupnos/crow/backend/internal/model"

type Channel struct {
	Name string `json:"name"`
	model.Model
}
