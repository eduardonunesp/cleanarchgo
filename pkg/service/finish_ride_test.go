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
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.Must(domain.RideCreate(domain.RideCreateParams{
		PassengerID: "2",
		InitialX:    "123",
		InitialY:    "321",
		FinalX:      "789",
		FinalY:      "987",
	})), nil)
	s.posRepo.EXPECT().GetPositionsByRideID("1").Return([]*domain.Position{
		domain.Must(domain.PositionRestore(domain.PositionRestoreParams{
			ID:     "1",
			RideID: "1",
			CoordX: "123",
			CoordY: "321",
			Date:   tNow,
		})),
		domain.Must(domain.PositionRestore(domain.PositionRestoreParams{
			ID:     "4",
			RideID: "1",
			CoordX: "789",
			CoordY: "987",
			Date:   tNow,
		})),
	}, nil)
	s.rideRepo.EXPECT().SaveRide(mock.Anything).Return(nil)
	err := s.useCase.Execute(FinishRideParams{
		RideID: "1",
	})
	s.NoError(err)
}
