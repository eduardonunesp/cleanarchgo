package service

import (
	"errors"
	"fmt"
)

type Error struct {
	WrappedErr error
}

func RaiseServiceError(err error) *Error {
	return &Error{err}
}

func ErrorIs(err, expectedErr error) bool {
	domainErr := new(Error)
	if ok := errors.As(err, &domainErr); !ok {
		return false
	}
	return errors.Is(domainErr.WrappedErr, expectedErr)
}

func (r *Error) Error() string {
	return fmt.Sprintf("service error: %s", r.WrappedErr.Error())
}
