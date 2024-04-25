package service

import (
	"testing"
	"time"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/test"
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
	s.rideRepo.EXPECT().GetRideByID("1").Return(domain.MustBuild(domain.RestoreRide(
		"1", "2", "3", "", "123", "321", "789", "987", "in_progress", tNow,
	)), nil)
	s.posRepo.EXPECT().GetPositionsByRideID("1").Return([]*domain.Position{
		domain.MustBuild(domain.BuildPosition(
			domain.PositionWithID("3"),
			domain.PositionWithRideID("1"),
			domain.PositionWithCoord("123", "321"),
			domain.PositionWithDate(tNow),
		)),
		domain.MustBuild(domain.BuildPosition(
			domain.PositionWithID("4"),
			domain.PositionWithRideID("1"),
			domain.PositionWithCoord("789", "987"),
			domain.PositionWithDate(tNow),
		)),
	}, nil)
	s.rideRepo.EXPECT().SaveRide(domain.MustBuild(domain.RestoreRide(
		"1", "2", "3", "10.00", "123", "321", "789", "987", "completed", tNow,
	))).Return(nil)
	err := s.useCase.Execute(&FinishRideParams{
		RideID: "1",
	})
	s.NoError(err)
}
