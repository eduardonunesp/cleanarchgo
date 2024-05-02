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
	position, err := PositionRestore(PositionRestoreParams{
		ID:     "1",
		RideID: "2",
		CoordX: "3",
		CoordY: "4",
		Date:   testTime,
	})
	s.NoError(err)
	s.Equal("1", position.ID().String())
	s.Equal("2", position.RideID().String())
	s.Equal("3", position.Coord().Lat())
	s.Equal("4", position.Coord().Long())
	s.Equal(valueobject.MustBuild(valueobject.DateFromUnix(testTime)), position.Date())
}
