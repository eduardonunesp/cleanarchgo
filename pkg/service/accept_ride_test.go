package service

import (
	"testing"
	"time"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestA(t *testing.T) {
	suite.Run(t, new(acceptRideTestSuite))
}

type acceptRideTestSuite struct {
	suite.Suite
	rideRepo *test.MockRideRepository
	accRepo  *test.MockAccountRepository
	useCase  *AcceptRide
}

func (s *acceptRideTestSuite) SetupTest() {
	s.rideRepo = test.NewMockRideRepository(s.T())
	s.accRepo = test.NewMockAccountRepository(s.T())
	s.useCase = NewAcceptRide(s.rideRepo, s.accRepo)
}

func (s *acceptRideTestSuite) TestAcceptRide() {
	tNow := time.Now().Unix()
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.Must(domain.BuildRide(
		domain.RideWithID("1"),
		domain.RideWithPassengerID("2"),
		domain.RideWithDriverID("3"),
		domain.RideWithFare("10.001"),
		domain.RideWithSegment(
			"-27.594870",
			"-48.548222",
			"-27.642040",
			"-48.669239",
		),
		domain.RideWithStatus("requested"),
		domain.RideWithDate(tNow),
	)), nil)
	s.accRepo.EXPECT().GetAccountByID("3").Return(domain.Must(domain.BuildAccount(
		domain.AccountWithID("3"),
		domain.AccountWithName("John Doe"),
		domain.AccountWithEmail("foo@bar.com"),
		domain.AccountWithCpf("11144477735"),
		domain.AccountWithCarPlate("ABC1234"),
		domain.AccountWithAccountType("driver"),
	)), nil)
	s.rideRepo.EXPECT().UpdateRide(mock.Anything).Return(nil)
	err := s.useCase.Execute(AcceptRideParams{
		RideID:   "1",
		DriverID: "3",
	})
	s.NoError(err)
}
