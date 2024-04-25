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
		PositionWithID("1"),
		PositionWithRideID("2"),
		PositionWithCoord("3", "4"),
		PositionWithDate(testTime),
	)
	s.NoError(err)
	s.Equal("1", position.ID().String())
	s.Equal("2", position.RideID().String())
	s.Equal("3", position.Coord().Lat())
	s.Equal("4", position.Coord().Long())
	s.Equal(valueobject.MustBuild(valueobject.DateFromUnix(testTime)), position.Date())
}
