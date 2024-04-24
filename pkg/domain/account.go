package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type Account struct {
	id          valueobject.UUID
	name        valueobject.Name
	email       valueobject.Email
	cpf         valueobject.Cpf
	carPlate    valueobject.CarPlate
	accountType valueobject.AccountType
}

func CreateAccount(name, email, cpf, carPlate, accountType string) (*Account, error) {
	var (
		newAcc Account
		err    error
	)
	newAcc.id = valueobject.MustUUID()
	if newAcc.name, err = valueobject.NameFromString(name); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.email, err = valueobject.EmailFromString(email); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.cpf, err = valueobject.CpfFromString(cpf); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.accountType, err = valueobject.AccountTypeFromString(accountType); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.accountType == valueobject.AccountTypeDriver {
		if newAcc.carPlate, err = valueobject.CarPlateFromString(carPlate); err != nil {
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
	if newAcc.id, err = valueobject.UUIDFromString(id); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.name, err = valueobject.NameFromString(name); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.email, err = valueobject.EmailFromString(email); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.cpf, err = valueobject.CpfFromString(cpf); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.accountType, err = valueobject.AccountTypeFromString(accountType); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newAcc.accountType == valueobject.AccountTypeDriver {
		if newAcc.carPlate, err = valueobject.CarPlateFromString(carPlate); err != nil {
			return nil, RaiseDomainError(err)
		}
	}
	return &newAcc, nil
}

func (a Account) ID() valueobject.UUID {
	return a.id
}

func (a Account) Name() valueobject.Name {
	return a.name
}

func (a Account) Email() valueobject.Email {
	return a.email
}

func (a Account) Cpf() valueobject.Cpf {
	return a.cpf
}

func (a Account) CarPlate() valueobject.CarPlate {
	return a.carPlate
}

func (a Account) AccountType() valueobject.AccountType {
	return a.accountType
}

func (a Account) IsPassenger() bool {
	return a.accountType == valueobject.AccountTypePassenger
}

func (a Account) IsDriver() bool {
	return a.accountType == valueobject.AccountTypeDriver
}
