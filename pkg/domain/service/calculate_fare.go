package service

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
)

func CalculateFare(_ []*domain.Position) (string, error) {
	return "10.00", nil
}
