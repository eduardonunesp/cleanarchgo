package domain

import (
	"fmt"

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

func (a Account) IsPassenger() bool {
	return a.accountType == valueobject.AccountTypePassenger
}

func (a Account) IsDriver() bool {
	return a.accountType == valueobject.AccountTypeDriver
}

func (a *Account) SetCarPlateOnce(carPlate string) error {
	var err error
	a.carPlate, err = valueobject.CarPlateFromString(carPlate)
	if err != nil {
		fmt.Println("err", err)
		return RaiseDomainError(err)
	}
	return nil
}

func AccountWithID(id string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.id, err = valueobject.UUIDFromString(id)
		return err
	}
}

func AccountWithName(name string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.name, err = valueobject.NameFromString(name)
		return err
	}
}

func AccountWithEmail(email string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.email, err = valueobject.EmailFromString(email)
		return err
	}
}

func AccountWithCpf(cpf string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.cpf, err = valueobject.CpfFromString(cpf)
		return err
	}
}

func AccountWithCarPlate(carPlate string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.carPlate, err = valueobject.CarPlateFromString(carPlate)
		return err
	}
}

func AccountWithHash(hash string) accountOption {
	return func(opt *Account) error {
		opt.hash = valueobject.LoadHashFromString(hash)
		return nil
	}
}

func AccountWithPassword(password string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.hash, err = valueobject.BuildHashFromString(password, nil)
		return err
	}
}

func AccountWithAccountType(accountType string) accountOption {
	return func(opt *Account) error {
		var err error
		opt.accountType, err = valueobject.AccountTypeFromString(accountType)
		return err
	}
}

func BuildAccount(options ...accountOption) (*Account, error) {
	newAcc := Account{
		id: valueobject.MustUUID(),
	}
	for _, opt := range options {
		if err := opt(&newAcc); err != nil {
			return nil, RaiseDomainError(err)
		}
	}
	return &newAcc, nil
}
