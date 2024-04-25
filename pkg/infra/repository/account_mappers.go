package repository

import (
	"fmt"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/db"
)

func mapDBAccountToDomainAccount(account *db.Account) (*domain.Account, error) {
	domainAcc, err := domain.RestoreAccount(
		fromPgTypeUUIDToString(account.ID),
		account.Name,
		account.Email,
		account.Cpf,
		account.CarPlate.String,
		string(account.AccountType.AccountType),
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
	pgTypeUUID, err := mapStringToPgTypeUUID(account.ID().String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse account uuid")
	}
	var accountType db.NullAccountType
	if err := accountType.Scan(account.AccountType().String()); err != nil {
		return nil, fmt.Errorf("failed to parse account type: %w", err)
	}
	return &db.SaveAccountParams{
		ID:          pgTypeUUID,
		Name:        account.Name().String(),
		Email:       account.Email().String(),
		Cpf:         account.Cpf().String(),
		CarPlate:    fromStringToPgTypeText(account.CarPlate().String()),
		AccountType: accountType,
	}, nil
}
