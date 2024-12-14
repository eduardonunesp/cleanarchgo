package service

import (
	"testing"
	"time"

	"github.com/eduardonunesp/cleanarchgo/internal/domain"
	"github.com/eduardonunesp/cleanarchgo/internal/test"
	"github.com/stretchr/testify/suite"
)

func TestRide(t *testing.T) {
	suite.Run(t, new(testRideSuite))
}

type testRideSuite struct {
	suite.Suite
	rideRepo *test.MockRideRepository
	accRepo  *test.MockAccountRepository
	useCase  *GetRide
}

func (s *testRideSuite) SetupTest() {
	s.rideRepo = test.NewMockRideRepository(s.T())
	s.accRepo = test.NewMockAccountRepository(s.T())
	s.useCase = NewGetRide(s.rideRepo, s.accRepo)
}

func (s *testRideSuite) TestGetRide() {
	tNow := time.Now().Unix()
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.Must(domain.RideRestore(domain.RideRestoreParams{
		ID:          "1",
		PassengerID: "2",
		DriverID:    "3",
		Fare:        "10.00",
		InitialX:    "-27.594870",
		InitialY:    "-48.548222",
		FinalX:      "-27.642040",
		FinalY:      "-48.669239",
		Status:      "requested",
		Date:        tNow,
	})), nil)
	s.accRepo.EXPECT().GetAccountByID("2").Return(domain.Must(domain.AccountRestore(domain.AccountRestoreParams{
		ID:           "2",
		Name:         "Foo Bar",
		Email:        "foobar@gmail.com",
		Cpf:          "11144477735",
		CarPlate:     "AAA9999",
		AccountType:  "passenger",
		PasswordHash: "",
		ConfirmedAt:  tNow,
		UpdatedAt:    tNow,
		CreatedAt:    tNow,
	})), nil)
	result, err := s.useCase.Execute(GetRideParams{
		RideID: "1",
	})
	s.NoError(err)
	s.NotNil(result)
	s.Equal(&GetRideResult{
		ID:             "1",
		PassengerID:    "2",
		FromLat:        "-27.594870",
		FromLong:       "-48.548222",
		ToLat:          "-27.642040",
		ToLong:         "-48.669239",
		Status:         "requested",
		PassengerName:  "Foo Bar",
		PassengerEmail: "foobar@gmail.com",
	}, result)
}
