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
	)
	s.NoError(err)
	s.Equal(MustBuildAccount(
		AccountWithID("1"),
		AccountWithName("Foo Bar"),
		AccountWithEmail("foo@bar.com"),
	), acc)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidName() {
	_, err := BuildAccount(
		AccountWithName("Foo"),
	)
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidName)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidEmail() {
	_, err := BuildAccount(
		AccountWithEmail("foo.com"),
	)
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidEmail)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCPF() {
	_, err := BuildAccount(
		AccountWithCPF("11144477700"),
	)
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidCPF)
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCarPlate() {
	_, err := BuildAccount(
		AccountIsDriver(),
		AccountWithCarPlate("AAA", true),
	)
	domainErr := new(Error)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, valueobject.ErrInvalidCarPlate)
}
