package domain

import (
	"testing"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
	"github.com/stretchr/testify/suite"
)

func mustBuildVO[T any](vo T, err error) T {
	if err != nil {
		panic(err)
	}
	return vo
}

func TestAccount(t *testing.T) {
	suite.Run(t, new(testAccountSuite))
}

type testAccountSuite struct {
	suite.Suite
}

func (s *testAccountSuite) TestBuildAccountWithSuccess() {
	acc, err := RestoreAccount("1", "Foo Bar", "foo@bar.com", "11144477735", "", valueobject.AccountTypePassenger.String())
	s.NoError(err)
	s.Equal(&Account{
		ID:          mustBuildVO(valueobject.UUIDFromString("1")),
		Name:        mustBuildVO(valueobject.NameFromString("Foo Bar")),
		Email:       mustBuildVO(valueobject.EmailFromString("foo@bar.com")),
		CPF:         mustBuildVO(valueobject.CpfFromString("11144477735")),
		AccountType: mustBuildVO(valueobject.AccountTypeFromString("passenger")),
	}, acc)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidName() {
	_, err := RestoreAccount("1", "", "foo@bar.com", "11144477735", "", valueobject.AccountTypeDriver.String())
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidName)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidEmail() {
	_, err := RestoreAccount("1", "Foo Bar", "foocom", "11144477735", "", valueobject.AccountTypeDriver.String())
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidEmail)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCPF() {
	_, err := RestoreAccount("1", "Foo Bar", "foo@bar.com", "11177735", "", valueobject.AccountTypeDriver.String())
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidCPF)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCarPlate() {
	_, err := RestoreAccount("1", "Foo Bar", "foo@bar.com", "11144477735", "AAA", valueobject.AccountTypeDriver.String())
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidCarPlate)
}
