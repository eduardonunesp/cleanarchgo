package domain

import (
	"testing"
	"time"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
	"github.com/stretchr/testify/suite"
)

func TestPosition(t *testing.T) {
	suite.Run(t, new(positionTestSuite))
}

type positionTestSuite struct {
	suite.Suite
}

func (s *positionTestSuite) TestBuildPosition() {
	testTime := time.Now().Unix()
	position, err := BuildPosition(
		WithPositionID("1"),
		WithRideID("2"),
		WithLatLong("3", "4"),
		WithDate(testTime),
	)
	s.NoError(err)
	s.Equal("1", string(position.PositionID))
	s.Equal("2", string(position.RideID))
	s.Equal("3", position.Coord.Lat)
	s.Equal("4", position.Coord.Long)
	s.Equal(valueobject.DateFromInt64(testTime), position.Date)
}
