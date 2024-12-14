package service

import (
	"testing"
	"time"

	"github.com/eduardonunesp/cleanarchgo/internal/domain"
	"github.com/eduardonunesp/cleanarchgo/internal/test"
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
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.Must(domain.RideRestore(domain.RideRestoreParams{
		ID:          "1",
		PassengerID: "2",
		DriverID:    "3",
		Fare:        "10.001",
		InitialX:    "-27.594870",
		InitialY:    "-48.548222",
		FinalX:      "-27.642040",
		FinalY:      "-48.669239",
		Status:      "requested",
		Date:        tNow,
	})), nil)
	s.accRepo.EXPECT().GetAccountByID("3").Return(domain.Must(domain.AccountRestore(domain.AccountRestoreParams{
		ID:           "3",
		Name:         "John Doe",
		Email:        "foo@bar.com",
		Cpf:          "11144477735",
		CarPlate:     "ABC1234",
		AccountType:  "driver",
		PasswordHash: "a",
		ConfirmedAt:  tNow,
		CreatedAt:    tNow,
		UpdatedAt:    tNow,
	})), nil)
	s.rideRepo.EXPECT().UpdateRide(mock.Anything).Return(nil)
	err := s.useCase.Execute(AcceptRideParams{
		RideID:   "1",
		DriverID: "3",
	})
	s.NoError(err)
}
