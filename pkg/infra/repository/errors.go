package repository

type RepositoryError struct {
	Err error
}

func RaiseRepositoryError(err error) *RepositoryError {
	return &RepositoryError{err}
}

func (r RepositoryError) Error() string {
	return r.Err.Error()
}
