package domain

func MustBuild[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}