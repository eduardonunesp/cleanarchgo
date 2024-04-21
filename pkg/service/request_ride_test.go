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
	s.accRepo.EXPECT().GetAccountByID("123").Return(&domain.Account{
		ID:          "123",
		Name:        "Foo Bar",
		Email:       "foobar@gmail.com",
		CPF:         "11144477735",
		IsPassenger: true,
	}, nil)
	s.rideRepo.EXPECT().HasActiveRideByPassengerID("123").Return(false, nil)
	s.rideRepo.EXPECT().SaveRide(mock.MatchedBy(func(ride *domain.Ride) bool {
		if ride.PassengerID != "123" {
			return false
		}
		if ride.FromLat != "123" {
			return false
		}
		if ride.FromLong != "321" {
			return false
		}
		if ride.ToLat != "789" {
			return false
		}
		if ride.ToLong != "987" {
			return false
		}
		return true
	})).Return(nil)

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
