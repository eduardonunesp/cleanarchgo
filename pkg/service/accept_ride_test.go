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
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.MustBuild(domain.RestoreRide(
		"1", "2", "3", "", "123", "321", "789", "987", "requested", tNow,
	)), nil)
	s.accRepo.EXPECT().GetAccountByID("3").Return(domain.MustBuildAccount(
		domain.AccountWithID("3"),
		domain.AccountWithName("John Doe"),
		domain.AccountWithEmail("foo@bar.com"),
		domain.AccountWithCpf("11144477735"),
		domain.AccountWithCarPlate("ABC1234"),
		domain.AccountWithAccountType("driver"),
	), nil)
	s.rideRepo.EXPECT().UpdateRide(mock.Anything).Return(nil)
	err := s.useCase.Execute(&AcceptRideParams{
		RideID:   "1",
		DriverID: "3",
	})
	s.NoError(err)
}
