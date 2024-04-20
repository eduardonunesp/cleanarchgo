package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestRide(t *testing.T) {
	suite.Run(t, new(testRideSuite))
}

type testRideSuite struct {
	suite.Suite
}

func (s *testRideSuite) TestBuildRideWithSuccess() {
	tNow := time.Now()
	rideStatus, err := BuildRideStatusFromString("requested")
	s.NoError(err)
	ride, err := BuildRide(
		RideWithIDAndPassengerID("1", "2"),
		RideWithDate(tNow),
	)
	s.NoError(err)
	s.Equal(&Ride{
		ID:          "1",
		PassengerID: "2",
		Date:        tNow,
		Status:      rideStatus,
	}, ride)
}

func (s *testRideSuite) TestBuildRideFailedInvalidID() {
	ride, err := BuildRide(
		RideWithIDAndPassengerID("", "2"),
	)
	domainErr := new(DomainError)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, ErrRideEmptyID)
	s.Nil(ride)
}

func (s *testRideSuite) TestBuildRideFailedInvalidPassengerID() {
	ride, err := BuildRide(
		RideWithIDAndPassengerID("1", ""),
	)
	domainErr := new(DomainError)
	s.ErrorAs(err, &domainErr)
	s.ErrorIs(domainErr.Err, ErrRideEmptyPassengerID)
	s.Nil(ride)
}
