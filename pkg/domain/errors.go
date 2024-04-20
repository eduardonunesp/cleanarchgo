package domain

type DomainError struct {
	Err error
}

func RaiseDomainError(err error) *DomainError {
	return &DomainError{err}
}

func (r DomainError) Error() string {
	return r.Err.Error()
}
