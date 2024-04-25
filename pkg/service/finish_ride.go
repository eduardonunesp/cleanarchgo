package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain/service"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	ErrFinishRideInvalidStatus = errors.New("invalid ride status")
)

type FinishRideParams struct {
	RideID string
}

type FinishRide struct {
	rideRepo repository.RideRepository
	posRepo  repository.PositionRepository
}

func NewFinishRide(
	rideRepo repository.RideRepository,
	posRepo repository.PositionRepository,
) *FinishRide {
	return &FinishRide{rideRepo, posRepo}
}

func (s FinishRide) Execute(input FinishRideParams) error {
	ride, err := s.rideRepo.GetRideByID(input.RideID)
	if err != nil {
		return err
	}
	if !ride.IsInProgress() {
		return ErrFinishRideInvalidStatus
	}
	positions, err := s.posRepo.GetPositionsByRideID(input.RideID)
	if err != nil {
		return err
	}
	fareToPay, err := service.CalculateFare(positions)
	if err != nil {
		return err
	}
	if err := ride.Finish(fareToPay); err != nil {
		return err
	}
	return s.rideRepo.SaveRide(ride)
}
