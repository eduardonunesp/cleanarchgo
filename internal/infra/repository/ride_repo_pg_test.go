//go:build db_integration

package repository

import (
	"testing"

	"github.com/eduardonunesp/cleanarchgo/internal/domain"
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
	domainRide := domain.Must(domain.RideCreate(domain.RideCreateParams{
		PassengerID: s.passengerUUID,
		InitialX:    "-27.594870",
		InitialY:    "-48.548222",
		FinalX:      "-27.642040",
		FinalY:      "-48.669239",
	}))
	err := s.rideDB.SaveRide(domainRide)
	s.NoError(err)
}

func (s *testRideRepoPgSuite) TestGetRideWithSuccess() {
	domainRide := domain.Must(domain.RideCreate(domain.RideCreateParams{
		PassengerID: s.passengerUUID,
		InitialX:    "-27.594870",
		InitialY:    "-48.548222",
		FinalX:      "-27.642040",
		FinalY:      "-48.669239",
	}))
	err := s.rideDB.SaveRide(domainRide)
	s.NoError(err)

	ride, err := s.rideDB.GetRideByID(domainRide.ID().String())
	s.NoError(err)
	s.Equal(domainRide.PassengerID(), ride.PassengerID())
	s.Equal("-27.594870", ride.Segment().From().Lat())
	s.Equal("-48.548222", ride.Segment().From().Long())
	s.Equal("-27.642040", ride.Segment().To().Lat())
	s.Equal("-48.669239", ride.Segment().To().Long())
}

func (s *testRideRepoPgSuite) TestUpdateRideWithSuccess() {
	domainAccount := domain.Must(domain.AccountCreate(domain.AccountCreateParams{
		Name:        "Foo Bar",
		Email:       "foo@bar.com",
		CPF:         "11144477735",
		Password:    "123456",
		AccountType: "passenger",
	}))

	domainRide := domain.Must(domain.RideCreate(domain.RideCreateParams{
		PassengerID: domainAccount.ID().String(),
		InitialX:    "-27.594870",
		InitialY:    "-48.548222",
		FinalX:      "-27.642040",
		FinalY:      "-48.669239",
	}))
	err := s.rideDB.SaveRide(domainRide)
	s.NoError(err)

	s.NoError(domainRide.Accept((domainAccount.ID())))

	err = s.rideDB.UpdateRide(domainRide)
	s.NoError(err)

	ride, err := s.rideDB.GetRideByID(domainRide.ID().String())
	s.NoError(err)
	s.Equal(domainRide.PassengerID(), ride.PassengerID())
	s.Equal(domainRide.DriverID(), ride.DriverID())
	s.Equal(domainRide.Status().String(), "accepted")
	s.Equal("-27.594870", ride.Segment().From().Lat())
	s.Equal("-48.548222", ride.Segment().From().Long())
	s.Equal("-27.642040", ride.Segment().To().Lat())
	s.Equal("-48.669239", ride.Segment().To().Long())
}
