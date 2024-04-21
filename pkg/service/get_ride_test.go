package service

import (
	"testing"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/test"
	"github.com/stretchr/testify/suite"
)

func TestRide(t *testing.T) {
	suite.Run(t, new(testRideSuite))
}

type testRideSuite struct {
	suite.Suite
	rideRepo    *test.MockRideRepository
	accountRepo *test.MockAccountRepository
	useCase     *GetRide
}

func (s *testRideSuite) SetupTest() {
	s.rideRepo = test.NewMockRideRepository(s.T())
	s.useCase = NewGetRide(s.rideRepo, s.accountRepo)
}

func (s *testRideSuite) TestGetRide() {
	s.rideRepo.EXPECT().GetRideByID("1").Return(&domain.Ride{
		ID: "1",
	}, nil)
	result, err := s.useCase.Execute(&GetRideParams{
		RideID: "1",
	})
	s.NoError(err)
	s.NotNil(result)
	s.Equal(&GetRideResult{
		ID: "1",
	}, result)
}
