//go:build db_integration

package repository

import (
	"testing"
	"time"

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
	domainRide, _ := domain.BuildRide(
		domain.RideWithID(s.uuid),
		domain.RideWithPassengerID(s.passengerUUID),
		domain.RideWithDriverID(s.driverUUID),
		domain.RideWithFare("10.00"),
	)
	err := s.rideDB.SaveRide(domainRide)
	s.NoError(err)
}

func (s *testRideRepoPgSuite) TestGetRideWithSuccess() {
	tNow := time.Now().UTC()
	domainRide, _ := domain.BuildRide(
		domain.RideWithID(s.uuid),
		domain.RideWithPassengerID(s.passengerUUID),
		domain.RideWithDriverID(s.driverUUID),
		domain.RideWithFare("10.001"),
		domain.RideWithDate(tNow),
		domain.RideWithDistance("500.00"),
		domain.RideWithFromLatLong("-27.594870", "-48.548222"),
		domain.RideWithToLatLong("-27.642040", "-48.669239"),
	)
	err := s.rideDB.SaveRide(domainRide)
	s.NoError(err)

	ride, err := s.rideDB.GetRideByID(s.uuid)
	s.NoError(err)
	s.Equal(s.uuid, ride.ID)
	s.Equal(s.passengerUUID, ride.PassengerID)
	s.Equal(s.driverUUID, ride.DriverID)
	s.Equal("10.001", ride.Fare)
	s.Equal("500.00", ride.Distance)
	s.Equal("-27.594870", ride.FromLat)
	s.Equal("-48.548222", ride.FromLong)
	s.Equal("-27.642040", ride.ToLat)
	s.Equal("-48.669239", ride.ToLong)
	s.Equal(tNow, ride.Date.UTC())
}
