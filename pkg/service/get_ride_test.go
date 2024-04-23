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
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.MustBuildRide(
		domain.RideWithID("1"),
		domain.RideWithPassengerID("2"),
		domain.RideWithStatus("requested"),
		domain.RideWithFromLatLongToLatLong("123", "321", "789", "987"),
	), nil)
	s.accRepo.EXPECT().GetAccountByID("2").Return(domain.MustBuildAccount(
		domain.AccountWithID("2"),
		domain.AccountWithName("Foo Bar"),
		domain.AccountWithEmail("foobar@gmail.com"),
	), nil)
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
