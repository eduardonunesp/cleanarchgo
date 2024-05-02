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
	acc, err := AccountCreate(AccountCreateParams{
		Name:        "Foo Bar",
		Email:       "foo@bar.com",
		CPF:         "11144477735",
		AccountType: "passenger",
		Password:    "supersecret",
	})
	s.NoError(err)
	s.Equal(acc.Name().String(), "Foo Bar")
	s.Equal(acc.Email().String(), "foo@bar.com")
	s.Equal(acc.Cpf().String(), "11144477735")
	s.Equal(acc.AccountType().String(), "passenger")
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidName() {
	_, err := AccountCreate(AccountCreateParams{
		Name: "",
	})
	s.True(ErrorIs(err, valueobject.ErrInvalidName), "Error should be invalid name")
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidEmail() {
	_, err := AccountCreate(AccountCreateParams{
		Name:  "Foo Bar",
		Email: "foo@",
	})
	s.True(ErrorIs(err, valueobject.ErrInvalidEmail), "Error should be invalid email")
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCPF() {
	_, err := AccountCreate(AccountCreateParams{
		Name:  "Foo Bar",
		Email: "foo@bar.com",
		CPF:   "11144477",
	})
	s.True(ErrorIs(err, valueobject.ErrInvalidCPF), "Error should be invalid cpf")
}

func (s *testAccountSuite) TestBuildAccountFailedInvalidCarPlate() {
	_, err := AccountCreate(AccountCreateParams{
		Name:        "Foo Bar",
		Email:       "foo@barcom.com",
		CPF:         "11144477735",
		AccountType: "driver",
		Password:    "supersecret",
		CarPlate:    "1234",
	})
	s.True(ErrorIs(err, valueobject.ErrInvalidCarPlate), "Error should be invalid car plate")
}

func (s *testAccountSuite) TestAccountCreatedAtAndUpdatedAt() {
	acc, err := AccountCreate(AccountCreateParams{
		Name:        "Foo Bar",
		Email:       "foo@bar.com",
		CPF:         "11144477735",
		AccountType: "passenger",
		Password:    "supersecret",
	})
	s.NoError(err)

	createdAt := acc.CreatedAt()
	updatedAt := acc.UpdatedAt()

	s.NotZero(createdAt, "Created at should not be zero")
	s.NotZero(updatedAt, "Updated at should not be zero")
}

func (s *testAccountSuite) TestAccountCreatedAt() {
	acc, err := AccountCreate(AccountCreateParams{
		Name:        "Foo Bar",
		Email:       "foo@bar.com",
		CPF:         "11144477735",
		AccountType: "passenger",
		Password:    "supersecret",
	})
	s.NoError(err)
	createdAt := acc.CreatedAt()
	s.NotZero(createdAt, "Created at should not be zero")
}
