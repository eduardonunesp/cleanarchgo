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
	rsRequested, err := BuildRideStatus("requested")
	s.NoError(err)
	s.Equal(rsRequested.String(), "requested")
	rsAccepted, err := BuildRideStatus("accepted")
	s.NoError(err)
	s.Equal(rsAccepted.String(), "accepted")
	rsInProgress, err := BuildRideStatus("in_progress")
	s.NoError(err)
	s.Equal(rsInProgress.String(), "in_progress")
}

func (s *testRideStatusSuite) TestRideStatusesFailed() {
	_, err := BuildRideStatus("non_sense_string")
	s.Error(err)
	_, noErr := BuildRideStatus("accepted")
	s.NoError(noErr)
}

func (s *testRideStatusSuite) TestRideStatusesFailedWithDomainError() {
	_, err := BuildRideStatus("non_sense_string")
	s.Error(err)
}
