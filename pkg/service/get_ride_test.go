package service

import (
	"testing"
	"time"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/test"
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
		InitialX:    "123",
		InitialY:    "321",
		FinalX:      "789",
		FinalY:      "987",
		Status:      "requested",
		Date:        tNow,
	})), nil)
	s.accRepo.EXPECT().GetAccountByID("2").Return(domain.Must(domain.AccountRestore(domain.AccountRestoreParams{
		ID:          "2",
		Name:        "Foo Bar",
		Email:       "foobar@gmail.com",
		Cpf:         "11144477735",
		CarPlate:    "AAA9999",
		AccountType: "passenger",
	})), nil)
	result, err := s.useCase.Execute(GetRideParams{
		RideID: "1",
	})
	s.NoError(err)
	s.NotNil(result)
	s.Equal(&GetRideResult{
		ID:             "1",
		PassengerID:    "2",
		FromLat:        "123",
		FromLong:       "321",
		ToLat:          "789",
		ToLong:         "987",
		Status:         "requested",
		PassengerName:  "Foo Bar",
		PassengerEmail: "foobar@gmail.com",
	}, result)
}
