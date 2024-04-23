package valueobject

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidCarPlate = errors.New("invalid car plate")
	checkCarPlateRE    = regexp.MustCompile(`[A-Z]{3}[0-9]{4}`)
)

type CarPlate string

func CarPlateFromString(s string) (CarPlate, error) {
	if !checkCarPlateRE.MatchString(s) {
		return "", ErrInvalidCarPlate
	}
	return CarPlate(s), nil
}

func (c CarPlate) String() string {
	return string(c)
}
