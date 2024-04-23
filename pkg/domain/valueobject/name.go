package valueobject

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidName = errors.New("invalid name")
	checkNameRE    = regexp.MustCompile(`[a-zA-Z] [a-zA-Z]+`)
)

type Name string

func NameFromString(s string) (Name, error) {
	if !checkNameRE.MatchString(s) {
		return "", ErrInvalidName
	}
	return Name(s), nil
}

func (n Name) String() string {
	return string(n)
}
