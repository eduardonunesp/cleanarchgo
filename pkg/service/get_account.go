package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errGetAccountNotFound = errors.New("account not found for given id")
)

type GetAccountInput struct {
	ID string
}

type GetAccountOuput struct {
	ID          string
	Name        string
	Email       string
	CPF         string
	CarPlate    string
	AccountType string
}

type GetAccount struct {
	accountRepo repository.AccountRepository
}

func NewGetAccount(accountRepo repository.AccountRepository) *GetAccount {
	return &GetAccount{accountRepo}
}

func (g GetAccount) Execute(input *GetAccountInput) (*GetAccountOuput, error) {
	account, err := g.accountRepo.GetAccountByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, RaiseServiceError(errGetAccountNotFound)
	}
	return &GetAccountOuput{
		ID:          account.ID.String(),
		Name:        account.Name.String(),
		Email:       account.Email.String(),
		CPF:         account.CPF.String(),
		CarPlate:    account.CarPlate.String(),
		AccountType: account.AccountType.String(),
	}, nil
}
