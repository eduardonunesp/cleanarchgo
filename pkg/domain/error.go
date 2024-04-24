package domain

import "fmt"

type Error struct {
	Err error
}

func RaiseDomainError(err error) *Error {
	return &Error{err}
}

func (r Error) Error() string {
	return fmt.Sprintf("domain error: %s", r.Err.Error())
}
