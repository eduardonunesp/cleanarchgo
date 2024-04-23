package valueobject

import "github.com/google/uuid"

type UUID string

func MustUUID() UUID {
	return UUID(uuid.Must(uuid.NewRandom()).String())
}

func UUIDFromString(s string) UUID {
	return UUID(s)
}

func (u UUID) String() string {
	return string(u)
}
