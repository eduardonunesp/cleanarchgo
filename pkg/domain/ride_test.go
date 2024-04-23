package domain

import (
	"testing"
	"time"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
	"github.com/stretchr/testify/suite"
)

func TestRide(t *testing.T) {
	suite.Run(t, new(testRideSuite))
}

type testRideSuite struct {
	suite.Suite
}

func (s *testRideSuite) TestBuildRideWithSuccess() {
	tNow := time.Now().Unix()
	rideStatus, err := valueobject.RideStatusFromString("requested")
	s.NoError(err)
	ride, err := BuildRide(
		RideWithID("1"),
		RideWithPassengerID("2"),
		RideWithDate(tNow),
	)
	s.NoError(err)
	s.Equal(&Ride{
		ID:          "1",
		PassengerID: "2",
		Date:        valueobject.DateFromInt64(tNow),
		Status:      rideStatus,
	}, ride)
}

func (s *testRideSuite) TestBuildRideFailedInvalidID() {
	ride, err := BuildRide(
		RideWithID(""),
		RideWithPassengerID("2"),
	)
	s.Error(err)
	s.Nil(ride)
}

func (s *testRideSuite) TestBuildRideFailedInvalidPassengerID() {
	ride, err := BuildRide(
		RideWithID("1"),
		RideWithPassengerID(""),
	)
	s.Error(err)
	s.Nil(ride)
}
