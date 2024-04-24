package service

import (
	"fmt"
)

type ServiceError struct {
	Err error
}

func RaiseServiceError(err error) *ServiceError {
	return &ServiceError{err}
}

func (r *ServiceError) Error() string {
	return fmt.Sprintf("service error: %s", r.Err.Error())
}
