package service

import (
	"testing"
	"time"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestFinishRide(t *testing.T) {
	suite.Run(t, new(testFinishRideSuite))
}

type testFinishRideSuite struct {
	suite.Suite
	rideRepo *test.MockRideRepository
	posRepo  *test.MockPositionRepository
	useCase  *FinishRide
}

func (s *testFinishRideSuite) SetupTest() {
	s.posRepo = test.NewMockPositionRepository(s.T())
	s.rideRepo = test.NewMockRideRepository(s.T())
	s.useCase = NewFinishRide(s.rideRepo, s.posRepo)
}

func (s *testFinishRideSuite) TestFinishRideSuccess() {
	tNow := time.Now().Unix()
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.Must(domain.RideRestore(domain.RideRestoreParams{
		ID:          "1",
		DriverID:    "3",
		PassengerID: "2",
		Fare:        "10.001",
		InitialX:    "-27.594870",
		InitialY:    "-48.548222",
		FinalX:      "-27.642040",
		FinalY:      "-48.669239",
		Status:      "in_progress",
		Date:        tNow,
	})), nil)
	s.posRepo.EXPECT().GetPositionsByRideID("1").Return([]*domain.Position{
		domain.Must(domain.PositionRestore(domain.PositionRestoreParams{
			ID:     "1",
			RideID: "1",
			CoordX: "-27.594870",
			CoordY: "-48.548222",
			Date:   tNow,
		})),
		domain.Must(domain.PositionRestore(domain.PositionRestoreParams{
			ID:     "4",
			RideID: "1",
			CoordX: "-27.642040",
			CoordY: "-48.669239",
			Date:   tNow,
		})),
	}, nil)
	s.rideRepo.EXPECT().SaveRide(mock.Anything).Return(nil)
	err := s.useCase.Execute(FinishRideParams{
		RideID: "1",
	})
	s.NoError(err)
}
