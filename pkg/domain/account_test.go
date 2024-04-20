package domain

import (
	"testing"

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
	)
	s.NoError(err)
	s.Equal(&Account{
		ID:    "1",
		Name:  "Foo Bar",
		Email: "foo@bar.com",
	}, acc)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidName() {
	_, err := BuildAccount(
		AccountWithName("Foo"),
	)
	domainErr := new(DomainError)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, ErrAccountInvalidName)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidEmail() {
	_, err := BuildAccount(
		AccountWithEmail("foo.com"),
	)
	domainErr := new(DomainError)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, ErrAccountInvalidEmail)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCPF() {
	_, err := BuildAccount(
		AccountWithCPF("11144477700"),
	)
	domainErr := new(DomainError)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, ErrAccountInvalidCPF)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCarPlate() {
	_, err := BuildAccount(
		AccountIsDriver(),
		AccountWithCarPlate("AAA"),
	)
	domainErr := new(DomainError)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, ErrAccountInvalidCarPlate)
}
