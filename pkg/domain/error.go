package domain

type Error struct {
	Err error
}

func RaiseDomainError(err error) *Error {
	return &Error{err}
}

func (r Error) Error() string {
	return r.Err.Error()
}
