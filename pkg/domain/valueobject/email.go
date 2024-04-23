package valueobject

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
	checkEmailRE    = regexp.MustCompile(`^(.+)@(.+)$`)
)

type Email struct {
	value string
}

func EmailFromString(s string) (Email, error) {
	var email Email
	if !checkEmailRE.MatchString(s) {
		return email, ErrInvalidEmail
	}
	email.value = s
	return email, nil
}

func (e Email) String() string {
	return string(e.value)
}
