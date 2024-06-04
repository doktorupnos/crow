package user

import (
	"github.com/google/uuid"
)

type Repo interface {
	Create(User) (uuid.UUID, error)
	GetAll() ([]User, error)
	GetByName(string) (User, error)
	GetByID(uuid.UUID) (User, error)
	Update(User) error
	Delete(User) error
}
