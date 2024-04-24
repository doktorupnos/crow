package user

import (
	"net/http"
	"regexp"

	"github.com/doktorupnos/crow/backend/internal/follow"
	"github.com/doktorupnos/crow/backend/internal/pages"
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

func (s *Service) Follow(u User, id uuid.UUID) error {
	o, err := s.GetByID(id)
	if err != nil {
		return err
	}

	return s.r.Follow(u, o)
}

func (s *Service) Unfollow(u User, id uuid.UUID) error {
	o, err := s.GetByID(id)
	if err != nil {
		return err
	}

	return s.r.Unfollow(u, o)
}

func (s *Service) Following(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]follow.Follow, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}

	return s.r.Following(LoadParams{
		UserID: userID,
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
	})
}

type LoadParams struct {
	UserID uuid.UUID
	pages.PaginationParams
}

func (s *Service) Followers(
	r *http.Request,
	defaultPageSize int,
	userID uuid.UUID,
) ([]follow.Follow, error) {
	page := pages.ExtractPage(r)
	limit := pages.ExtractLimit(r)
	if limit == 0 {
		limit = defaultPageSize
	}

	return s.r.Followers(LoadParams{
		UserID: userID,
		PaginationParams: pages.PaginationParams{
			PageNumber: page,
			PageSize:   limit,
		},
	})
}

func (s *Service) FollowingCount(u User) (int, error) {
	return s.r.FollowingCount(u)
}

func (s *Service) FollowerCount(u User) (int, error) {
	return s.r.FollowersCount(u)
}

func (s *Service) FollowsUser(u, t User) (bool, error) {
	return s.r.FollowsUser(u, t)
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
