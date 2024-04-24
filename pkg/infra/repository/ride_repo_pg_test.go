//go:build db_integration

package repository

import (
	"testing"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestRideRepoDB(t *testing.T) {
	suite.Run(t, new(testRideRepoPgSuite))
}

type testRideRepoPgSuite struct {
	suite.Suite
	rideDB        *RideRepositoryPG
	uuid          string
	passengerUUID string
	driverUUID    string
}

func (s *testRideRepoPgSuite) SetupTest() {
	s.rideDB = NewRideRepositoryPG("postgres://postgres:123456@localhost:5432/postgres?sslmode=disable")
	s.uuid = uuid.Must(uuid.NewRandom()).String()
	s.driverUUID = uuid.Must(uuid.NewRandom()).String()
	s.passengerUUID = uuid.Must(uuid.NewRandom()).String()
}

func (s *testRideRepoPgSuite) TestCreateRideWithSuccess() {
	domainRide, _ := domain.CreateRide(
		s.passengerUUID,
		"10.001",
		"-48.669239",
		"-27.594870",
		"-48.548222",
		"-27.642040",
	)
	err := s.rideDB.SaveRide(domainRide)
	s.NoError(err)
}

func (s *testRideRepoPgSuite) TestGetRideWithSuccess() {
	domainRide, err := domain.CreateRide(
		s.passengerUUID,
		"10.001",
		"-27.594870",
		"-48.548222",
		"-27.642040",
		"-48.669239",
	)
	s.NoError(err)
	err = s.rideDB.SaveRide(domainRide)
	s.NoError(err)

	ride, err := s.rideDB.GetRideByID(domainRide.ID.String())
	s.NoError(err)
	s.Equal(domainRide.PassengerID, ride.PassengerID)
	s.Equal(domainRide.DriverID, ride.DriverID)
	s.Equal("10.001", ride.Fare)
	s.Equal("-27.594870", ride.Segment.From().Lat())
	s.Equal("-48.548222", ride.Segment.From().Long())
	s.Equal("-27.642040", ride.Segment.To().Lat())
	s.Equal("-48.669239", ride.Segment.To().Long())
}
