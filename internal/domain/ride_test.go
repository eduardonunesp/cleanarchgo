package domain

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestRide(t *testing.T) {
	suite.Run(t, new(testRideSuite))
}

type testRideSuite struct {
	suite.Suite
}

func (s *testRideSuite) TestBuildRideWithSuccess() {
	tmpUUID := "1580134f-0853-4aef-a9d8-245389912900"
	ride, err := RideCreate(RideCreateParams{
		PassengerID: tmpUUID,
		InitialX:    "1",
		InitialY:    "2",
		FinalX:      "3",
		FinalY:      "4",
	})
	s.NoError(err)
	s.Equal(tmpUUID, ride.PassengerID().String())
	s.Equal("1", ride.Segment().From().Lat())
	s.Equal("2", ride.Segment().From().Long())
	s.Equal("3", ride.Segment().To().Lat())
	s.Equal("4", ride.Segment().To().Long())
	s.Equal("requested", ride.Status().String())
}

func (s *testRideSuite) TestBuildRideFailedInvalidID() {
	ride, err := RideRestore(RideRestoreParams{
		PassengerID: "1",
	})
	s.Error(err)
	s.Nil(ride)
}

func (s *testRideSuite) TestBuildRideFailedInvalidPassengerID() {
	ride, err := RideRestore(RideRestoreParams{
		ID: "1",
	})
	s.Error(err)
	s.Nil(ride)
}
