package user

import "regexp"

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrNameEmpty     Error = "name is empty"
	ErrNameSmall     Error = "name is too small"
	ErrNameBig       Error = "name is too big"
	ErrNameMalformed Error = "name malformed"
	ErrNameTaken     Error = "name taken"

	ErrPasswordEmpty     Error = "password is empty"
	ErrPasswordBig       Error = "password is too big"
	ErrPasswordMalformed Error = "password malformed"

	nameMinLen     = 3
	nameMaxLen     = 20
	passwordMaxLen = 72
)

var (
	usernameRE = regexp.MustCompile(`^[a-zA-Z0-9_.]+$`)
	passwordRE = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&* \-]+$`)
)

type CreateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r CreateRequest) Validate() error {
	if err := validateName(r.Name); err != nil {
		return err
	}
	if err := validatePassword(r.Password); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	switch l := len(name); {
	case l == 0:
		return ErrNameEmpty
	case l < nameMinLen:
		return ErrNameSmall
	case l > nameMaxLen:
		return ErrNameBig
	case !usernameRE.MatchString(name):
		return ErrNameMalformed
	}
	return nil
}

func validatePassword(password string) error {
	switch l := len(password); {
	case l == 0:
		return ErrPasswordEmpty
	case l > passwordMaxLen:
		return ErrPasswordBig
	case !passwordRE.MatchString(password):
		return ErrPasswordMalformed
	}
	return nil
}
