package valueobject

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmptyUUID = errors.New("uuid cannot be empty")
)

type UUID string

func MustUUID() UUID {
	return UUID(uuid.Must(uuid.NewRandom()).String())
}

func UUIDFromString(s string) (UUID, error) {
	if s == "" {
		return "", ErrEmptyUUID
	}
	return UUID(s), nil
}

func (u UUID) String() string {
	return string(u)
}
