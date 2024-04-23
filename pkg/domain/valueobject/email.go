package valueobject

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
	checkEmailRE    = regexp.MustCompile(`^(.+)@(.+)$`)
)

type Email string

func EmailFromString(s string) (Email, error) {
	if !checkEmailRE.MatchString(s) {
		return "", ErrInvalidEmail
	}
	return Email(s), nil
}

func (e Email) String() string {
	return string(e)
}
