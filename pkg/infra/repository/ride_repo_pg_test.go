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
}

func (s *testRideRepoPgSuite) SetupTest() {
	s.rideDB = NewRideRepositoryPG("postgres://postgres:123456@localhost:5432/postgres?sslmode=disable")
	s.uuid = uuid.Must(uuid.NewRandom()).String()
	s.passengerUUID = uuid.Must(uuid.NewRandom()).String()
}

func (s *testRideRepoPgSuite) TestCreateRideWithSuccess() {
	domainRide, _ := domain.BuildRide(
		domain.RideWithIDAndPassengerID(s.uuid, s.passengerUUID),
	)
	err := s.rideDB.SaveRide(domainRide)
	s.NoError(err)
}
