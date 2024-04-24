package follow

import "github.com/google/uuid"

type Follow struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}
