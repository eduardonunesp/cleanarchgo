package service

import (
	"testing"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/gateway"
	"github.com/eduardonunesp/cleanarchgo/pkg/test"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestSignup(t *testing.T) {
	suite.Run(t, new(testSignupSuite))
}

type testSignupSuite struct {
	suite.Suite
	accountRepo *test.MockAccountRepository
	useCase     *Signup
}

func (s *testSignupSuite) SetupTest() {
	s.accountRepo = test.NewMockAccountRepository(s.T())
	s.useCase = NewSignup(s.accountRepo, gateway.NewMailerGatewayMemory())
}

func (s *testSignupSuite) TestSignupSuccess() {
	s.accountRepo.EXPECT().HasAccountByEmail("foobar@gmail.com").Return(false, nil)
	s.accountRepo.EXPECT().SaveAccount(mock.MatchedBy(func(acc *domain.Account) bool {
		if acc.Name != "Foo Bar" {
			return false
		}
		if acc.Email != "foobar@gmail.com" {
			return false
		}
		if acc.CPF != "11144477735" {
			return false
		}
		if acc.ID == "" {
			return false
		}
		return true
	})).Return(nil)
	result, err := s.useCase.Execute(&SignupParams{
		Name:        "Foo Bar",
		Email:       "foobar@gmail.com",
		CPF:         "11144477735",
		IsPassenger: true,
	})
	s.NoError(err)
	s.NotNil(result)
}

func (s *testSignupSuite) TestSignupFailedAccountExists() {
	s.accountRepo.EXPECT().HasAccountByEmail("foobar@gmail.com").Return(true, nil)
	result, err := s.useCase.Execute(&SignupParams{
		Name:        "Foo Bar",
		Email:       "foobar@gmail.com",
		CPF:         "11144477735",
		IsPassenger: true,
	})
	serviceErr := new(ServiceError)
	s.ErrorAs(err, &serviceErr)
	s.Nil(result)
}
