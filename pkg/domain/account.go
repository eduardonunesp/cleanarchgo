package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type accountOption func(*Account) error

type Account struct {
	id          valueobject.UUID
	name        valueobject.Name
	email       valueobject.Email
	cpf         valueobject.Cpf
	carPlate    valueobject.CarPlate
	hash        valueobject.Hash
	accountType valueobject.AccountType
	confirmedAt valueobject.Date
	createdAt   valueobject.Date
	updatedAt   valueobject.Date
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

func (a Account) Hash() valueobject.Hash {
	return a.hash
}

func (a Account) CarPlate() valueobject.CarPlate {
	return a.carPlate
}

func (a Account) AccountType() valueobject.AccountType {
	return a.accountType
}

func (a Account) IsConfirmed() bool {
	return !a.confirmedAt.IsZero()
}

func (a Account) CreatedAt() valueobject.Date {
	return a.createdAt
}

func (a Account) UpdatedAt() valueobject.Date {
	return a.updatedAt
}

func (a Account) IsPassenger() bool {
	return a.accountType == valueobject.AccountTypePassenger
}

func (a Account) IsDriver() bool {
	return a.accountType == valueobject.AccountTypeDriver
}

func (a *Account) SetCarPlate(carPlate string) error {
	var err error
	a.carPlate, err = valueobject.CarPlateFromString(carPlate)
	if err != nil {
		return RaiseDomainError(err)
	}
	return nil
}

func (a *Account) Confirm() {
	a.confirmedAt = valueobject.DateFromNow()
}
