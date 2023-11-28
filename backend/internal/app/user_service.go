package app

import (
	"regexp"

	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type UserService struct {
	ur user.UserRepo
}

func NewUserService(ur user.UserRepo) *UserService {
	return &UserService{ur}
}

func (s *UserService) Create(name, password string) error {
	if err := validateName(name); err != nil {
		return err
	}
	if err := validatePassword(password); err != nil {
		return err
	}

	hashed, err := hashPassword(password)
	if err != nil {
		return err
	}

	return s.ur.Create(user.User{Name: name, Password: hashed})
}

func (s *UserService) GetAll() ([]user.User, error) {
	return s.ur.GetAll()
}

func (s *UserService) GetByName(name string) (user.User, error) {
	return s.ur.GetByName(name)
}

func (s *UserService) GetByID(id uuid.UUID) (user.User, error) {
	return s.ur.GetByID(id)
}

func (s *UserService) Update(u user.User, name, password string) error {
	if err := validateName(name); err != nil {
		return err
	}
	if err := validatePassword(password); err != nil {
		return err
	}

	hashed, err := hashPassword(password)
	if err != nil {
		return err
	}

	u.Name = name
	u.Password = hashed
	return s.ur.Update(u)
}

func (s *UserService) Delete(u user.User) error {
	return s.ur.Delete(u)
}

type ErrUser string

func (e ErrUser) Error() string {
	return string(e)
}

const (
	ErrUserNameEmpty         = ErrUser("name is empty")
	ErrUserNameTooBig        = ErrUser("name is too big")
	ErrUserNameMalformed     = ErrUser("malformed name")
	ErrUserPasswordEmpty     = ErrUser("password is empty")
	ErrUserPasswordTooBig    = ErrUser("password too big")
	ErrUserPasswordMalformed = ErrUser("malformed password")
)

const pattern = "^[a-zA-Z0-9_]+$"

var userRegex = regexp.MustCompile(pattern)

func validateName(name string) error {
	if len(name) == 0 {
		return ErrUserNameEmpty
	}
	if len(name) > 20 {
		return ErrUserNameTooBig
	}

	if !userRegex.MatchString(name) {
		return ErrUserNameMalformed
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) == 0 {
		return ErrUserPasswordEmpty
	}

	if len(password) > 72 {
		return ErrUserPasswordTooBig
	}

	if !userRegex.MatchString(password) {
		return ErrUserPasswordMalformed
	}

	return nil
}
