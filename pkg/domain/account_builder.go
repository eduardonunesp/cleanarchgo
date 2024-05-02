package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type AccountCreateParams struct {
	Name        string
	Email       string
	CPF         string
	Password    string
	CarPlate    string
	AccountType string
}

type AccountRestoreParams struct {
	ID           string
	Name         string
	Email        string
	Cpf          string
	CarPlate     string
	AccountType  string
	PasswordHash string
	ConfirmedAt  int64
	CreatedAt    int64
	UpdatedAt    int64
}

func AccountCreate(params AccountCreateParams) (*Account, error) {
	var (
		newAccount Account
		err        error
	)
	newAccount.id = valueobject.MustUUID()
	if newAccount.name, err = valueobject.NameFromString(params.Name); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid name")
	}
	if newAccount.email, err = valueobject.EmailFromString(params.Email); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid email")
	}
	if newAccount.cpf, err = valueobject.CpfFromString(params.CPF); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid cpf")
	}
	if newAccount.hash, err = valueobject.BuildHashFromString(params.Password, nil); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid password")
	}
	if newAccount.accountType, err = valueobject.AccountTypeFromString(params.AccountType); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid account type")
	}
	if newAccount.accountType == valueobject.AccountTypeDriver {
		if newAccount.carPlate, err = valueobject.CarPlateFromString(params.CarPlate); err != nil {
			return nil, RaiseDomainErrorWith(err, "invalid car plate")
		}
	}
	newAccount.createdAt = valueobject.DateFromNow()
	newAccount.updatedAt = valueobject.DateFromNow()
	return &newAccount, err
}

func AccountRestore(params AccountRestoreParams) (*Account, error) {
	var (
		account Account
		err     error
	)
	if account.id, err = valueobject.UUIDFromString(params.ID); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid id")
	}
	if account.name, err = valueobject.NameFromString(params.Name); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid name")
	}
	if account.email, err = valueobject.EmailFromString(params.Email); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid email")
	}
	if account.cpf, err = valueobject.CpfFromString(params.Cpf); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid cpf")
	}
	if account.hash, err = valueobject.LoadHashFromString(params.PasswordHash); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid password hash")
	}
	if account.accountType, err = valueobject.AccountTypeFromString(params.AccountType); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid account type")
	}
	if account.accountType == valueobject.AccountTypeDriver {
		if account.carPlate, err = valueobject.CarPlateFromString(params.CarPlate); err != nil {
			return nil, RaiseDomainErrorWith(err, "invalid car plate")
		}
	}
	if params.ConfirmedAt != 0 {
		if account.confirmedAt, err = valueobject.DateFromUnix(params.ConfirmedAt); err != nil {
			return nil, RaiseDomainErrorWith(err, "invalid confirmed at")
		}
	}
	if account.createdAt, err = valueobject.DateFromUnix(params.CreatedAt); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid created at")
	}
	if account.updatedAt, err = valueobject.DateFromUnix(params.UpdatedAt); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid updated at")
	}
	return &account, err
}
