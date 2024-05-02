package repository

type Error struct {
	Err error
}

func RaiseRepositoryError(err error) *Error {
	return &Error{err}
}

func (r Error) Error() string {
	return r.Err.Error()
}
