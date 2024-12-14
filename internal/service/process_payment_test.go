package service

import (
	"testing"
	"time"

	"github.com/eduardonunesp/cleanarchgo/internal/domain"
	"github.com/eduardonunesp/cleanarchgo/internal/test"
	"github.com/stretchr/testify/suite"
)

func TestProcessPayment(t *testing.T) {
	suite.Run(t, new(testProcessPaymentSuite))
}

type testProcessPaymentSuite struct {
	suite.Suite
	// Mocks goes here
	rideRepo *test.MockRideRepository
	ccGW     *test.MockCreditCardGW
	useCase  *ProcessPayment
}

func (s *testProcessPaymentSuite) SetupTest() {
	// Mocks goes here
	s.rideRepo = test.NewMockRideRepository(s.T())
	s.ccGW = test.NewMockCreditCardGW(s.T())
	s.useCase = NewProcessPayment(s.rideRepo, s.ccGW)
}

func (s *testProcessPaymentSuite) TestProcessPaymentSuccess() {
	tNow := time.Now().Unix()
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.Must(domain.RideRestore(domain.RideRestoreParams{
		ID:          "1",
		PassengerID: "2",
		DriverID:    "3",
		Fare:        "10.00",
		Status:      "completed",
		Date:        tNow,
		InitialX:    "-27.594870",
		InitialY:    "-48.548222",
		FinalX:      "-27.642040",
		FinalY:      "-48.669239",
	})), nil)
	s.ccGW.EXPECT().ProcessPayment("123", "10.00").Return(nil)
	err := s.useCase.Execute(ProcessPaymentParams{
		RideID:          "1",
		CreditCardToken: "123",
		Amount:          "10.00",
	})
	s.NoError(err)
}
