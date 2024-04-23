package domain

import (
	"errors"
	"fmt"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type AccountOption func(acc *Account) error

type Account struct {
	ID          valueobject.UUID
	Name        valueobject.Name
	Email       valueobject.Email
	CPF         valueobject.Cpf
	CarPlate    valueobject.CarPlate
	Available   bool
	IsDriver    bool
	IsPassenger bool
}

func AccountWithID(id string) AccountOption {
	return func(acc *Account) error {
		var err error
		if acc.ID, err = valueobject.UUIDFromString(id); err != nil {
			return err
		}
		return nil
	}
}

func AccountWithName(name string) AccountOption {
	return func(acc *Account) error {
		var err error
		if acc.Name, err = valueobject.NameFromString(name); err != nil {
			return err
		}
		return nil
	}
}

func AccountWithEmail(email string) AccountOption {
	return func(acc *Account) error {
		var err error
		if acc.Email, err = valueobject.EmailFromString(email); err != nil {
			return err
		}
		return nil
	}
}

func AccountWithCarPlate(carPlate string) AccountOption {
	return func(acc *Account) error {
		if !acc.IsDriver {
			return nil
		}
		var err error
		if acc.CarPlate, err = valueobject.CarPlateFromString(carPlate); err != nil {
			return err
		}
		return nil
	}
}

func AccountWithCPF(cpf string) AccountOption {
	return func(acc *Account) error {
		var err error
		if acc.CPF, err = valueobject.CpfFromString(cpf); err != nil {
			return err
		}
		return nil
	}
}

func AccountIsDriver() AccountOption {
	return func(acc *Account) error {
		acc.IsDriver = true
		return nil
	}
}

func AccountSetDriver(value bool) AccountOption {
	return func(acc *Account) error {
		acc.IsDriver = value
		return nil
	}
}

func AccountSetPassenger(value bool) AccountOption {
	return func(acc *Account) error {
		acc.IsPassenger = value
		return nil
	}
}

func AccountIsPassenger() AccountOption {
	return func(acc *Account) error {
		acc.IsPassenger = true
		return nil
	}
}

func AccountIsDriverAvailable() AccountOption {
	return func(acc *Account) error {
		acc.Available = true
		return nil
	}
}

func AccountSetAvailable(value bool) AccountOption {
	return func(acc *Account) error {
		if !acc.IsDriver {
			return errors.New("only drivers can be available for rides")
		}
		acc.Available = value
		return nil
	}
}

func BuildAccount(accOpts ...AccountOption) (*Account, error) {
	var newAcc Account
	for _, accOpt := range accOpts {
		if accOpt == nil {
			continue
		}
		if err := accOpt(&newAcc); err != nil {
			return nil, RaiseDomainError(fmt.Errorf("failed to build Account: %w", err))
		}
	}
	if newAcc.ID.String() == "" {
		newAcc.ID = valueobject.MustUUID()
	}
	return &newAcc, nil
}

func MustBuildAccount(accOpts ...AccountOption) *Account {
	acc, err := BuildAccount(accOpts...)
	if err != nil {
		panic(err)
	}
	return acc
}

func (a Account) AuthorizeDriver() error {
	if !a.IsDriver {
		return RaiseDomainError(fmt.Errorf("account %s is not a driver", a.ID))
	}
	if !a.Available {
		return RaiseDomainError(fmt.Errorf("driver %s is not available", a.ID))
	}
	return nil
}
