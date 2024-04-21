package domain

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

var (
	errAccountInvalidEmail    = errors.New("invalid email")
	errAccountInvalidName     = errors.New("invalid name")
	errAccountInvalidCPF      = errors.New("invalid CPF")
	errAccountInvalidCarPlate = errors.New("invalid car plate")
)

var (
	checkNameRE     = regexp.MustCompile(`[a-zA-Z] [a-zA-Z]+`)
	checkEmailRE    = regexp.MustCompile(`^(.+)@(.+)$`)
	checkCarPlateRE = regexp.MustCompile(`[A-Z]{3}[0-9]{4}`)
)

type AccountOption func(acc *Account) error

type Account struct {
	ID          string
	Name        string
	Email       string
	CPF         string
	CarPlate    string
	IsDriver    bool
	IsPassenger bool
}

func AccountWithID(id string) AccountOption {
	return func(acc *Account) error {
		acc.ID = id
		return nil
	}
}

func AccountWithName(name string) AccountOption {
	return func(acc *Account) error {
		if !checkNameRE.MatchString(name) {
			return errAccountInvalidName
		}
		acc.Name = name
		return nil
	}
}

func AccountWithEmail(email string) AccountOption {
	return func(acc *Account) error {
		if !checkEmailRE.MatchString(email) {
			return errAccountInvalidEmail
		}
		acc.Email = email
		return nil
	}
}

func AccountWithCarPlate(carPlate string) AccountOption {
	return func(acc *Account) error {
		if acc.IsDriver && !checkCarPlateRE.MatchString(carPlate) {
			return errAccountInvalidCarPlate
		}
		acc.CarPlate = carPlate
		return nil
	}
}

func AccountWithCPF(cpf string) AccountOption {
	return func(acc *Account) error {
		if !validate(cpf) {
			return errAccountInvalidCPF
		}
		acc.CPF = cpf
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

func BuildAccount(accOpts ...AccountOption) (*Account, error) {
	var newAcc Account
	for _, accOpt := range accOpts {
		if accOpt == nil {
			continue
		}
		if err := accOpt(&newAcc); err != nil {
			return nil, RaiseDomainError(fmt.Errorf("failed to build account: %w", err))
		}
	}
	if newAcc.ID == "" {
		newAcc.ID = uuid.Must(uuid.NewRandom()).String()
	}
	return &newAcc, nil
}
