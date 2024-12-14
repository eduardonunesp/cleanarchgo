package domain

import (
	"errors"
	"fmt"
)

type Error struct {
	WrappedErr error
}

func RaiseDomainError(err error) *Error {
	return &Error{err}
}

func RaiseDomainErrorWith(err error, msg string) *Error {
	return &Error{fmt.Errorf("%s: %w", msg, err)}
}

func ErrorIs(err, expectedErr error) bool {
	domainErr := new(Error)
	if ok := errors.As(err, &domainErr); !ok {
		return false
	}
	return errors.Is(domainErr.WrappedErr, expectedErr)
}

func (r *Error) Error() string {
	return fmt.Sprintf("domain error: %s", r.WrappedErr.Error())
}
