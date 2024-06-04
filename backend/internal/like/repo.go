package like

import "github.com/google/uuid"

type Repo interface {
	Create(Like) error
	Load() ([]Like, error)
	Single(userID, postID uuid.UUID) (Like, error)
	Delete(Like) error
}
