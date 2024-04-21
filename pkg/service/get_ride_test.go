package service

import (
	"testing"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/test"
	"github.com/stretchr/testify/suite"
)

func TestRide(t *testing.T) {
	suite.Run(t, new(testRideSuite))
}

type testRideSuite struct {
	suite.Suite
	rideRepo *test.MockRideRepository
	accRepo  *test.MockAccountRepository
	useCase  *GetRide
}

func (s *testRideSuite) SetupTest() {
	s.rideRepo = test.NewMockRideRepository(s.T())
	s.accRepo = test.NewMockAccountRepository(s.T())
	s.useCase = NewGetRide(s.rideRepo, s.accRepo)
}

func (s *testRideSuite) TestGetRide() {
	s.rideRepo.EXPECT().GetRideByID("1").Return(&domain.Ride{
		ID:          "1",
		PassengerID: "2",
		Status:      domain.RideStatusRequested,
		FromLat:     "123",
		FromLong:    "321",
		ToLat:       "789",
		ToLong:      "987",
	}, nil)
	s.accRepo.EXPECT().GetAccountByID("2").Return(&domain.Account{
		ID:    "2",
		Name:  "Foo Bar",
		Email: "foobar@gmail.com",
	}, nil)
	result, err := s.useCase.Execute(&GetRideParams{
		RideID: "1",
	})
	s.NoError(err)
	s.NotNil(result)
	s.Equal(&GetRideResult{
		ID:             "1",
		PassengerID:    "2",
		FromLat:        "123",
		FromLong:       "321",
		ToLat:          "789",
		ToLong:         "987",
		Status:         "requested",
		PassengerName:  "Foo Bar",
		PassengerEmail: "foobar@gmail.com",
	}, result)
}
