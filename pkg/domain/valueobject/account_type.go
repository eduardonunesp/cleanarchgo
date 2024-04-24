package valueobject

import "errors"

var ErrInvalidAccountType = errors.New("invalid account type")

type AccountType int

const (
	AccountTypeDriver AccountType = iota
	AccountTypePassenger
)

func AccountTypeFromString(s string) (AccountType, error) {
	switch s {
	case "driver":
		return AccountTypeDriver, nil
	case "passenger":
		return AccountTypePassenger, nil
	default:
		return 0, ErrInvalidAccountType
	}
}

func (a AccountType) String() string {
	switch a {
	case AccountTypeDriver:
		return "driver"
	case AccountTypePassenger:
		return "passenger"
	default:
		return ""
	}
}
