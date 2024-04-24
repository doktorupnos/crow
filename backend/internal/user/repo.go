package user

import (
	"github.com/google/uuid"
)

type Repo interface {
	Create(u User) (uuid.UUID, error)
	GetAll() ([]User, error)
	GetByName(name string) (User, error)
	GetByID(id uuid.UUID) (User, error)
	Update(u User) error
	Delete(u User) error
}
