package domain

import (
	"fmt"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type Account struct {
	ID          valueobject.UUID
	Name        valueobject.Name
	Email       valueobject.Email
	CPF         valueobject.Cpf
	CarPlate    valueobject.CarPlate
	AccountType valueobject.AccountType
}

func CreateAccount(name, email, cpf, carPlate, accountType string) (*Account, error) {
	var (
		newAcc Account
		err    error
	)
	newAcc.ID = valueobject.MustUUID()
	if newAcc.Name, err = valueobject.NameFromString(name); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.Email, err = valueobject.EmailFromString(email); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.CPF, err = valueobject.CpfFromString(cpf); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.AccountType, err = valueobject.AccountTypeFromString(accountType); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.AccountType == valueobject.AccountTypeDriver {
		if newAcc.CarPlate, err = valueobject.CarPlateFromString(carPlate); err != nil {
			return nil, RaiseDomainError(err)
		}
	}
	return &newAcc, nil
}

func RestoreAccount(id, name, email, cpf, carPlate, accountType string) (*Account, error) {
	var (
		newAcc Account
		err    error
	)
	if newAcc.ID, err = valueobject.UUIDFromString(id); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.Name, err = valueobject.NameFromString(name); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.Email, err = valueobject.EmailFromString(email); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.CPF, err = valueobject.CpfFromString(cpf); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.AccountType, err = valueobject.AccountTypeFromString(accountType); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.AccountType == valueobject.AccountTypeDriver {
		if newAcc.CarPlate, err = valueobject.CarPlateFromString(carPlate); err != nil {
			return nil, RaiseDomainError(err)
		}
	}
	return &newAcc, nil
}

func (a Account) IsPassenger() bool {
	return a.AccountType == valueobject.AccountTypePassenger
}

func (a Account) IsDriver() bool {
	return a.AccountType == valueobject.AccountTypeDriver
}

func (a Account) AuthorizeDriver() error {
	if a.AccountType != valueobject.AccountTypeDriver {
		return RaiseDomainError(fmt.Errorf("account %s is not a driver", a.ID))
	}
	return nil
}
