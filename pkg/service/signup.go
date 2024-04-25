package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/gateway"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errSignupAccountExists = errors.New("account already exists for this email")
)

type SignupParams struct {
	Name        string
	Email       string
	CPF         string
	CarPlate    string
	AccountType string
}

type SignupResult struct {
	AccountUID string
}

type Signup struct {
	accountRepo repository.AccountRepository
	mailerGW    gateway.MailerGW
}

func NewSignup(accountRepo repository.AccountRepository, mailerGW gateway.MailerGW) *Signup {
	return &Signup{accountRepo, mailerGW}
}

func (s Signup) Execute(input *SignupParams) (*SignupResult, error) {
	accountExists, err := s.accountRepo.HasAccountByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if accountExists {
		return nil, RaiseServiceError(errSignupAccountExists)
	}
	domainAccount, err := domain.BuildAccount(
		domain.AccountWithName(input.Name),
		domain.AccountWithEmail(input.Email),
		domain.AccountWithCpf(input.CPF),
		domain.AccountWithAccountType(input.AccountType),
	)
	if err != nil {
		return nil, err
	}
	if domainAccount.IsDriver() {
		if err := domainAccount.SetCarPlateOnce(input.CarPlate); err != nil {
			return nil, err
		}
	}
	if err := s.accountRepo.SaveAccount(domainAccount); err != nil {
		return nil, err
	}
	s.mailerGW.Send(input.Email, "Welcome!", "")
	return &SignupResult{
		domainAccount.ID().String(),
	}, nil
}
