package service

import (
	"testing"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/test"

	"github.com/stretchr/testify/suite"
)

func TestGetAccount(t *testing.T) {
	suite.Run(t, new(testGetAccountSuite))
}

type testGetAccountSuite struct {
	suite.Suite
	accountRepo *test.MockAccountRepository
	useCase     *GetAccount
}

func (s *testGetAccountSuite) SetupTest() {
	s.accountRepo = test.NewMockAccountRepository(s.T())
	s.useCase = NewGetAccount(s.accountRepo)
}

func (s *testGetAccountSuite) TestGetAccountSuccess() {
	s.accountRepo.EXPECT().GetAccountByID("1").Return(
		domain.MustBuild(domain.BuildAccount(
			domain.AccountWithID("1"),
			domain.AccountWithName("Foo Bar"),
			domain.AccountWithEmail("foo@bar.com"),
			domain.AccountWithCpf("11144477735"),
			domain.AccountWithCarPlate("AAA9999"),
			domain.AccountWithAccountType("driver"),
		)), nil)
	result, err := s.useCase.Execute(&GetAccountInput{
		ID: "1",
	})
	s.NoError(err)
	s.NotNil(result)
	s.Equal(&GetAccountOuput{
		ID:          "1",
		Name:        "Foo Bar",
		Email:       "foo@bar.com.br",
		CPF:         "11144477735",
		CarPlate:    "AAA9999",
		AccountType: "driver",
	}, result)
}

func (s *testGetAccountSuite) TestGetAccountFailedUserNotFound() {
	s.accountRepo.EXPECT().GetAccountByID("2").Return(nil, nil)
	result, err := s.useCase.Execute(&GetAccountInput{
		ID: "2",
	})
	s.Error(err)
	s.Nil(result)
}
