package valueobject

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidCarPlate = errors.New("invalid car plate")
	checkCarPlateRE    = regexp.MustCompile(`[A-Z]{3}[0-9]{4}`)
)

type CarPlate struct {
	value string
}

func CarPlateFromString(s string) (CarPlate, error) {
	var carPlate CarPlate
	if !checkCarPlateRE.MatchString(s) {
		return carPlate, ErrInvalidCarPlate
	}
	carPlate.value = s
	return carPlate, nil
}

func (c CarPlate) String() string {
	return string(c.value)
}
