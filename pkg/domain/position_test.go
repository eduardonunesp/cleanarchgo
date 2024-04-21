package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestPosition(t *testing.T) {
	suite.Run(t, new(positionTestSuite))
}

type positionTestSuite struct {
	suite.Suite
}

func (s *positionTestSuite) TestBuildPosition() {
	testTime := time.Now()
	position, err := BuildPosition(
		WithPositionID("1"),
		WithRideID("2"),
		WithLat("3"),
		WithLong("4"),
		WithDate(testTime),
	)
	s.NoError(err)
	s.Equal("1", position.PositionID)
	s.Equal("2", position.RideID)
	s.Equal("3", position.Lat)
	s.Equal("4", position.Long)
	s.Equal(testTime, position.Date)
}
