package service

import (
	"testing"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestRequestRide(t *testing.T) {
	suite.Run(t, new(testRequestRideSuite))
}

type testRequestRideSuite struct {
	suite.Suite
	accRepo  *test.MockAccountRepository
	rideRepo *test.MockRideRepository
	useCase  *RequestRide
}

func (s *testRequestRideSuite) SetupTest() {
	s.accRepo = test.NewMockAccountRepository(s.T())
	s.rideRepo = test.NewMockRideRepository(s.T())
	s.useCase = NewRequestRide(s.rideRepo, s.accRepo)
}

func (s *testRequestRideSuite) TestRequestRideSuccess() {
	s.accRepo.EXPECT().GetAccountByID("123").Return(domain.Must(domain.AccountRestore(domain.AccountRestoreParams{
		ID:           "123",
		Name:         "Foo Bar",
		Email:        "foobar@gmail.com",
		Cpf:          "11144477735",
		AccountType:  "passenger",
		PasswordHash: "123",
		CreatedAt:    1,
		UpdatedAt:    1,
	})), nil)
	s.rideRepo.EXPECT().HasActiveRideByPassengerID("123").Return(false, nil)
	s.rideRepo.EXPECT().SaveRide(mock.Anything).Return(nil)

	result, err := s.useCase.Execute(RequestRideParams{
		PassengerID: "123",
		FromLat:     "123",
		FromLong:    "321",
		ToLat:       "789",
		ToLong:      "987",
	})
	s.NoError(err)
	s.NotNil(result)
}
