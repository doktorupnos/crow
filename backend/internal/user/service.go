package user

import (
	"regexp"

	"github.com/doktorupnos/crow/backend/internal/passwd"
	"github.com/google/uuid"
)

type Service struct {
	r Repo
}

func NewService(r Repo) *Service {
	return &Service{r}
}

func (s *Service) Create(name, password string) (uuid.UUID, error) {
	if err := validateName(name); err != nil {
		return uuid.UUID{}, err
	}
	if err := validatePassword(password); err != nil {
		return uuid.UUID{}, err
	}

	hashed, err := passwd.Hash(password)
	if err != nil {
		return uuid.UUID{}, err
	}

	return s.r.Create(User{Name: name, Password: hashed})
}

func (s *Service) GetAll() ([]User, error) {
	return s.r.GetAll()
}

func (s *Service) GetByName(name string) (User, error) {
	return s.r.GetByName(name)
}

func (s *Service) GetByID(id uuid.UUID) (User, error) {
	return s.r.GetByID(id)
}

func (s *Service) Update(u User, name, password string) error {
	if err := validateName(name); err != nil {
		return err
	}
	if err := validatePassword(password); err != nil {
		return err
	}

	hashed, err := passwd.Hash(password)
	if err != nil {
		return err
	}

	u.Name = name
	u.Password = hashed
	return s.r.Update(u)
}

func (s *Service) Delete(u User) error {
	return s.r.Delete(u)
}

type ErrUser string

func (e ErrUser) Error() string {
	return string(e)
}

const (
	ErrUserNameEmpty         = ErrUser("name empty")
	ErrUserNameTooBig        = ErrUser("name too big")
	ErrUserNameTooSmall      = ErrUser("password too small")
	ErrUserNameMalformed     = ErrUser("name malformed")
	ErrUserNameTaken         = ErrUser("name taken")
	ErrUserPasswordEmpty     = ErrUser("password empty")
	ErrUserPasswordTooBig    = ErrUser("password too big")
	ErrUserPasswordMalformed = ErrUser("password malformed")
)

var (
	userRegex     = regexp.MustCompile("^[a-zA-Z0-9_.]+$")
	passwordRegex = regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*]")
)

func validateName(name string) error {
	if len(name) == 0 {
		return ErrUserNameEmpty
	}

	if len(name) < 4 {
		return ErrUserNameTooSmall
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

	if !passwordRegex.MatchString(password) {
		return ErrUserPasswordMalformed
	}

	return nil
}
