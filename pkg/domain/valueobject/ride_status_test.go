package valueobject

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestRideStatus(t *testing.T) {
	suite.Run(t, new(testRideStatusSuite))
}

type testRideStatusSuite struct {
	suite.Suite
}

func (s *testRideStatusSuite) TestRideStatusesSuccess() {
	rsRequested, err := RideStatusFromString("requested")
	s.NoError(err)
	s.Equal(rsRequested, RideStatusRequested)
	rsCompleted, err := RideStatusFromString("completed")
	s.NoError(err)
	s.Equal(rsCompleted, RideStatusCompleted)
	rsAccepted, err := RideStatusFromString("accepted")
	s.NoError(err)
	s.Equal(rsAccepted, RideStatusAccepted)
	rsInProgress, err := RideStatusFromString("in_progress")
	s.NoError(err)
	s.Equal(rsInProgress, RideStatusInProgres)
}

func (s *testRideStatusSuite) TestRideStatusesFailed() {
	_, err := RideStatusFromString("non_sense_string")
	s.Error(err)
	_, noErr := RideStatusFromString("completed")
	s.NoError(noErr)
}

func (s *testRideStatusSuite) TestRideStatusesFailedWithDomainError() {
	_, err := RideStatusFromString("non_sense_string")
	s.Error(err)
}
