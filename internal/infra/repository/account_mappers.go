package repository

import (
	"fmt"

	"github.com/eduardonunesp/cleanarchgo/internal/domain"
	"github.com/eduardonunesp/cleanarchgo/internal/infra/db"
)

func mapDBAccountToDomainAccount(account *db.Account) (*domain.Account, error) {
	return domain.AccountRestore(domain.AccountRestoreParams{
		ID:           fromPgTypeUUIDToString(account.ID),
		Name:         account.Name,
		Email:        account.Email,
		Cpf:          account.Cpf,
		CarPlate:     fromPgTypeTextToString(account.CarPlate),
		AccountType:  string(account.AccountType.AccountType),
		PasswordHash: account.Hashpw,
		ConfirmedAt:  account.Confirmedat.Time.Unix(),
		CreatedAt:    account.Confirmedat.Time.Unix(),
		UpdatedAt:    account.Confirmedat.Time.Unix(),
	})
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
