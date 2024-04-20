//go:build db_integration

package repository

import (
	"fmt"
	"testing"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestAccountRepoDB(t *testing.T) {
	suite.Run(t, new(testAccountRepoDBSuite))
}

type testAccountRepoDBSuite struct {
	suite.Suite
	accDB *AccountRepositoryPG
	uuid  string
}

func (s *testAccountRepoDBSuite) SetupTest() {
	s.accDB = NewAccountRepositoryPG("postgres://postgres:123456@localhost:5432/postgres?sslmode=disable")
	s.uuid = uuid.Must(uuid.NewRandom()).String()
}

func (s *testAccountRepoDBSuite) TestCreateAccountWithSuccess() {
	domainAcc, _ := domain.BuildAccount(
		domain.AccountWithName("Foo Bar"),
		domain.AccountWithEmail(fmt.Sprintf("foo%s@gmail.com", s.uuid)),
		domain.AccountWithCPF("11144477735"),
		domain.AccountIsPassenger(),
	)
	err := s.accDB.SaveAccount(domainAcc)
	s.NoError(err)
}

func (s *testAccountRepoDBSuite) TestGetAccountWithSuccess() {
	domainAcc, _ := domain.BuildAccount(
		domain.AccountWithName("Foo Bar"),
		domain.AccountWithEmail(fmt.Sprintf("foo%s@gmail.com", s.uuid)),
		domain.AccountWithCPF("11144477735"),
		domain.AccountIsPassenger(),
	)
	err := s.accDB.SaveAccount(domainAcc)
	s.NoError(err)

	acc, err := s.accDB.GetAccountByID(domainAcc.ID)
	s.NoError(err)
	s.NotNil(acc)
}

func (s *testAccountRepoDBSuite) TestCreateAccountFailedDuplicatedEmail() {
	email := fmt.Sprintf("foobar%s@gmail.com", s.uuid)
	domainAcc, _ := domain.BuildAccount(
		domain.AccountWithName("Foo Bar"),
		domain.AccountWithEmail(email),
		domain.AccountWithCPF("11144477735"),
		domain.AccountIsPassenger(),
	)
	err := s.accDB.SaveAccount(domainAcc)
	s.NoError(err)

	err = s.accDB.SaveAccount(domainAcc)
	s.Error(err)
}