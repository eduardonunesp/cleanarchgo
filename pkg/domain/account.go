package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type accountOption func(opt *Account) error

// type CreateAccountParams struct {
// 	Name        string
// 	Email       string
// 	Cpf         string
// 	CarPlate    string
// 	AccountType string
// }

// type ResstoreAccountParams struct {
// 	ID          string
// 	Name        string
// 	Email       string
// 	Cpf         string
// 	CarPlate    string
// 	AccountType string
// }

type Account struct {
	id          valueobject.UUID
	name        valueobject.Name
	email       valueobject.Email
	cpf         valueobject.Cpf
	carPlate    valueobject.CarPlate
	accountType valueobject.AccountType
}

// func CreateAccountFromParams(params *CreateAccountParams) (*Account, error) {
// 	return CreateAccount(params.Name, params.Email, params.Cpf, params.CarPlate, params.AccountType)
// }

// func MustCrateAccountFromParams(params *CreateAccountParams) *Account {
// 	acc, err := CreateAccountFromParams(params)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return acc
// }

// func RestoreAccountFromParams(params *ResstoreAccountParams) (*Account, error) {
// 	return RestoreAccount(params.ID, params.Name, params.Email, params.Cpf, params.CarPlate, params.AccountType)
// }

// func MustRestoreAccountFromParams(params *ResstoreAccountParams) *Account {
// 	acc, err := RestoreAccountFromParams(params)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return acc
// }

// func CreateAccount(name, email, cpf, carPlate, accountType string) (*Account, error) {
// 	var (
// 		newAcc Account
// 		err    error
// 	)
// 	newAcc.id = valueobject.MustUUID()
// 	if newAcc.name, err = valueobject.NameFromString(name); err != nil {
// 		return nil, RaiseDomainError(err)
// 	}
// 	if newAcc.email, err = valueobject.EmailFromString(email); err != nil {
// 		return nil, RaiseDomainError(err)
// 	}
// 	if newAcc.cpf, err = valueobject.CpfFromString(cpf); err != nil {
// 		return nil, RaiseDomainError(err)
// 	}
// 	if newAcc.accountType, err = valueobject.AccountTypeFromString(accountType); err != nil {
// 		return nil, RaiseDomainError(err)
// 	}
// 	if newAcc.accountType == valueobject.AccountTypeDriver {
// 		if newAcc.carPlate, err = valueobject.CarPlateFromString(carPlate); err != nil {
// 			return nil, RaiseDomainError(err)
// 		}
// 	}
// 	return &newAcc, nil
// }

// func RestoreAccount(id, name, email, cpf, carPlate, accountType string) (*Account, error) {
// 	var (
// 		newAcc Account
// 		err    error
// 	)
// 	if newAcc.id, err = valueobject.UUIDFromString(id); err != nil {
// 		return nil, RaiseDomainError(err)
// 	}
// 	if newAcc.name, err = valueobject.NameFromString(name); err != nil {
// 		return nil, RaiseDomainError(err)
// 	}
// 	if newAcc.email, err = valueobject.EmailFromString(email); err != nil {
// 		return nil, RaiseDomainError(err)
// 	}
// 	if newAcc.cpf, err = valueobject.CpfFromString(cpf); err != nil {
// 		return nil, RaiseDomainError(err)
// 	}
// 	if newAcc.accountType, err = valueobject.AccountTypeFromString(accountType); err != nil {
// 		return nil, RaiseDomainError(err)
// 	}
// 	if newAcc.accountType == valueobject.AccountTypeDriver {
// 		if newAcc.carPlate, err = valueobject.CarPlateFromString(carPlate); err != nil {
// 			return nil, RaiseDomainError(err)
// 		}
// 	}
// 	return &newAcc, nil
// }

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
	if newAcc.accountType == valueobject.AccountTypeDriver && newAcc.carPlate.String() == "" {
		return nil, RaiseDomainError(valueobject.ErrInvalidCarPlate)
	}
	return &newAcc, nil
}

func MustBuildAccount(options ...accountOption) *Account {
	acc, err := BuildAccount(options...)
	if err != nil {
		panic(err)
	}
	return acc
}
