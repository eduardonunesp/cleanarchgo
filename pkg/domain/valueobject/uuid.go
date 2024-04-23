package valueobject

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmptyUUID = errors.New("uuid cannot be empty")
)

type UUID struct {
	value string
}

func MustUUID() UUID {
	return UUID{uuid.Must(uuid.NewRandom()).String()}
}

func UUIDFromString(s string) (UUID, error) {
	var uuid UUID
	if s == "" {
		return uuid, ErrEmptyUUID
	}
	uuid.value = s
	return uuid, nil
}

func (u UUID) String() string {
	return u.value
}
