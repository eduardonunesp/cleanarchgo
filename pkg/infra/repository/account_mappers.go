package repository

import (
	"fmt"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/db"
)

func mapDBAccountToDomainAccount(account *db.Account) (*domain.Account, error) {
	domainAcc, err := domain.BuildAccount(
		domain.AccountWithID(fromPgTypeUUIDToString(account.ID)),
		domain.AccountWithName(account.Name),
		domain.AccountWithEmail(account.Email),
		domain.AccountWithCPF(account.Cpf),
		domain.AccountWithCarPlate(account.CarPlate.String, account.IsDriver),
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
	pgTypeUUID, err := mapStringToPgTypeUUID(account.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse account uuid")
	}
	return &db.SaveAccountParams{
		ID:          pgTypeUUID,
		Name:        account.Name.String(),
		Email:       account.Email.String(),
		Cpf:         account.CPF.String(),
		CarPlate:    fromStringToPgTypeText(account.CarPlate.String()),
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}, nil
}
