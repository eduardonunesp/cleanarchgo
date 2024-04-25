package domain

import (
	"testing"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
	"github.com/stretchr/testify/suite"
)

func TestAccount(t *testing.T) {
	suite.Run(t, new(testAccountSuite))
}

type testAccountSuite struct {
	suite.Suite
}

func (s *testAccountSuite) TestBuildAccountWithSuccess() {
	acc, err := BuildAccount(
		AccountWithID("1"),
		AccountWithName("Foo Bar"),
		AccountWithEmail("foo@bar.com"),
		AccountWithCpf("11144477735"),
		AccountWithAccountType("passenger"),
	)
	s.NoError(err)
	s.Equal(acc.ID().String(), "1")
	s.Equal(acc.Name().String(), "Foo Bar")
	s.Equal(acc.Email().String(), "foo@bar.com")
	s.Equal(acc.Cpf().String(), "11144477735")
	s.Equal(acc.AccountType().String(), "passenger")
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidName() {
	_, err := BuildAccount(
		AccountWithName(""),
	)
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidName)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidEmail() {
	_, err := BuildAccount(
		AccountWithEmail(""),
	)
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidEmail)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCPF() {
	_, err := BuildAccount(
		AccountWithCpf(""),
	)
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidCPF)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCarPlate() {
	_, err := BuildAccount(
		AccountWithCarPlate(""),
	)
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidCarPlate)
}
