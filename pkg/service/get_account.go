package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errGetAccountNotFound = errors.New("account not found for given id")
)

type GetAccountParams struct {
	ID string
}

type GetAccountResult struct {
	ID          string
	Name        string
	Email       string
	CPF         string
	CarPlate    string
	IsPassenger bool
	IsDriver    bool
}

type GetAccount struct {
	accountRepo repository.AccountRepository
}

func NewGetAccount(accountRepo repository.AccountRepository) *GetAccount {
	return &GetAccount{accountRepo}
}

func (g GetAccount) Execute(input *GetAccountParams) (*GetAccountResult, error) {
	account, err := g.accountRepo.GetAccountByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, RaiseServiceError(errGetAccountNotFound)
	}
	return &GetAccountResult{
		ID:          account.ID,
		Name:        account.Name,
		Email:       account.Email,
		CPF:         account.CPF,
		CarPlate:    account.CarPlate,
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
	}, nil
}
