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
	IsPassenger bool
	IsDriver    bool
}

type GetAccount struct {
	accountRepo repository.AccountRepository
}

func NewGetAccount(accountRepo repository.AccountRepository) *GetAccount {
	return &GetAccount{accountRepo}
}

func (g GetAccount) Execute(input *GetAccountInput) (*GetAccountOuput, error) {
	// account, err := g.accountRepo.GetAccountByID(input.ID)
	// if err != nil {
	// 	return nil, err
	// }
	// if account == nil {
	// 	return nil, RaiseServiceError(errGetAccountNotFound)
	// }
	// return &GetAccountOuput{
	// 	ID:          string(account.ID),
	// 	Name:        string(account.Name),
	// 	Email:       string(account.Email),
	// 	CPF:         string(account.CPF),
	// 	CarPlate:    string(account.CarPlate),
	// 	IsPassenger: account.IsPassenger,
	// 	IsDriver:    account.IsDriver,
	// }, nil
	return nil, nil
}
