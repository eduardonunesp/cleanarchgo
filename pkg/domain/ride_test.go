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
	tNow := time.Now().Unix()
	ride, err := BuildRide(
		RideWithID("1"),
		RideWithPassengerID("2"),
		RideWithDate(tNow),
		RideWithStatus("requested"),
	)
	s.NoError(err)
	s.Equal(MustBuildRide(
		RideWithID("1"),
		RideWithPassengerID("2"),
		RideWithDate(tNow),
		RideWithStatus("requested"),
	), ride)
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
