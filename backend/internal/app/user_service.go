package app

import (
	"net/http"
	"regexp"

	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/user"
	"github.com/google/uuid"
)

type UserService struct {
	ur user.UserRepo
}

func NewUserService(ur user.UserRepo) *UserService {
	return &UserService{ur}
}

func (s *UserService) Create(name, password string) (uuid.UUID, error) {
	if err := validateName(name); err != nil {
		return uuid.UUID{}, err
	}
	if err := validatePassword(password); err != nil {
		return uuid.UUID{}, err
	}

	hashed, err := hashPassword(password)
	if err != nil {
		return uuid.UUID{}, err
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

func (s *UserService) Follow(u user.User, id uuid.UUID) error {
	o, err := s.GetByID(id)
	if err != nil {
		return err
	}

	return s.ur.Follow(u, o)
}

func (s *UserService) Unfollow(u user.User, id uuid.UUID) error {
	o, err := s.GetByID(id)
	if err != nil {
		return err
	}

	return s.ur.Unfollow(u, o)
}

func (s *UserService) Following(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]user.Follow, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}

	return s.ur.Following(user.LoadParams{
		UserID: userID,
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
	})
}

func (s *UserService) Followers(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]user.Follow, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}

	return s.ur.Followers(user.LoadParams{
		UserID: userID,
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
	})
}

func (s *UserService) FollowingCount(u user.User) (int, error) {
	return s.ur.FollowingCount(u)
}

func (s *UserService) FollowerCount(u user.User) (int, error) {
	return s.ur.FollowersCount(u)
}

func (s *UserService) FollowsUser(u, t user.User) (bool, error) {
	return s.ur.FollowsUser(u, t)
}

type ErrUser string

func (e ErrUser) Error() string {
	return string(e)
}

const (
	ErrUserNameEmpty         = ErrUser("name empty")
	ErrUserNameTooBig        = ErrUser("name too big")
	ErrUserNameMalformed     = ErrUser("name malformed")
	ErrUserNameTaken         = ErrUser("name taken")
	ErrUserPasswordEmpty     = ErrUser("password empty")
	ErrUserPasswordTooBig    = ErrUser("password too big")
	ErrUserPasswordMalformed = ErrUser("password malformed")
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
