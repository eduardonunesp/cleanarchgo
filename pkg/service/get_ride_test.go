package service

import (
	"testing"
	"time"

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
	tNow := time.Now().Unix()
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.Must(domain.BuildRide(
		domain.RideWithID("1"),
		domain.RideWithPassengerID("2"),
		domain.RideWithDriverID("3"),
		domain.RideWithFare("10.00"),
		domain.RideWithSegment(
			"123",
			"321",
			"789",
			"987",
		),
		domain.RideWithStatus("requested"),
		domain.RideWithDate(tNow),
	)), nil)
	s.accRepo.EXPECT().GetAccountByID("2").Return(domain.Must(domain.BuildAccount(
		domain.AccountWithID("2"),
		domain.AccountWithName("Foo Bar"),
		domain.AccountWithEmail("foobar@gmail.com"),
		domain.AccountWithCpf("11144477735"),
		domain.AccountWithCarPlate("AAA9999"),
		domain.AccountWithAccountType("passenger"),
	)), nil)
	result, err := s.useCase.Execute(GetRideParams{
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
