package valueobject

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidName = errors.New("invalid name")
	checkNameRE    = regexp.MustCompile(`[a-zA-Z] [a-zA-Z]+`)
)

type Name struct {
	value string
}

func NameFromString(s string) (Name, error) {
	var name Name
	if !checkNameRE.MatchString(s) {
		return name, ErrInvalidName
	}
	name.value = s
	return name, nil
}

func (n Name) String() string {
	return n.value
}
