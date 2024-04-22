package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errRideNotInProgress = errors.New("ride status is not in progress")
)

type UpdatePositionParams struct {
	RideID string
	Lat    string
	Long   string
}

type UpdatePosition struct {
	rideRepo repository.RideRepository
	posRepo  repository.PositionRepository
}

func NewUpdatePosition(rideRepo repository.RideRepository, posRepo repository.PositionRepository) *UpdatePosition {
	return &UpdatePosition{rideRepo, posRepo}
}

func (s UpdatePosition) Execute(params *UpdatePositionParams) error {
	ride, err := s.rideRepo.GetRideByID(params.RideID)
	if err != nil {
		return err
	}
	if ride == nil {
		return RaiseServiceError(errAcceptRideNotFound)
	}
	if ride.Status != domain.RideStatusInProgres {
		return RaiseServiceError(errRideNotInProgress)
	}
	position, err := domain.BuildPosition(
		domain.WithRideID(params.RideID),
		domain.WithLat(params.Lat),
		domain.WithLong(params.Long),
	)
	if err != nil {
		return err
	}
	if err := s.posRepo.SavePosition(position); err != nil {
		return err
	}
	return nil
}
