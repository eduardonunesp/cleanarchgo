package domain

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
	rsRequested, err := BuildRideStatusFromString("requested")
	s.NoError(err)
	s.Equal(rsRequested, RideStatusRequested)

	rsCompleted, err := BuildRideStatusFromString("completed")
	s.NoError(err)
	s.Equal(rsCompleted, RideStatusCompleted)
}

func (s *testRideStatusSuite) TestRideStatusesFailed() {
	_, err := BuildRideStatusFromString("non_sense_string")
	domainErr := new(DomainError)
	s.ErrorAs(err, &domainErr)

	_, noErr := BuildRideStatusFromString("completed")
	s.NoError(noErr)
}
