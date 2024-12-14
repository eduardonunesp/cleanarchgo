package valueobject

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrNotValidUUIDV4 = errors.New("uuid is not a valid v4 uuid")
)

type UUID struct {
	value string
}

func MustUUID() UUID {
	return UUID{uuid.Must(uuid.NewRandom()).String()}
}

func UUIDFromString(s string) (UUID, error) {
	result, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, ErrNotValidUUIDV4
	}
	return UUID{
		value: result.String(),
	}, nil
}

func (u UUID) String() string {
	return u.value
}
