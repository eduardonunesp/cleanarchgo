package repository

import (
	"fmt"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/db"
)

func mapDBAccountToDomainAccount(account *db.Account) (*domain.Account, error) {
	if account == nil {
		return nil, fmt.Errorf("db account cannot be nil")
	}
	domainAcc, err := domain.BuildAccount(
		domain.AccountWithID(fromPgTypeUUIDToString(account.ID)),
		domain.AccountWithName(account.Name),
		domain.AccountWithEmail(account.Email),
		domain.AccountWithCPF(account.Cpf),
		domain.AccountWithCarPlate(account.CarPlate.String),
		domain.AccountSetDriver(account.IsDriver),
		domain.AccountSetPassenger(account.IsPassenger),
	)
	if err != nil {
		return nil, err
	}
	return domainAcc, nil
}

func mapDomainAccountToSaveAccountParams(account *domain.Account) (*db.SaveAccountParams, error) {
	if account == nil {
		return nil, fmt.Errorf("domain account cannot be nil")
	}
	pgTypeUUID, err := mapStringToPgTypeUUID(account.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse account uuid")
	}
	return &db.SaveAccountParams{
		ID:          pgTypeUUID,
		Name:        account.Name,
		Email:       account.Email,
		Cpf:         account.CPF,
		CarPlate:    fromStringToPgTypeText(account.CarPlate),
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}, nil
}
