package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errStartRideInvalidStatus = errors.New("ride status is not accepted")
)

type StartRideParams struct {
	RideID string
}

type StartRide struct {
	rideRepo repository.RideRepository
}

func NewStartRide(rideRepo repository.RideRepository) *StartRide {
	return &StartRide{rideRepo}
}

func (s StartRide) Execute(params *StartRideParams) error {
	// ride, err := s.rideRepo.GetRideByID(params.RideID)
	// if err != nil {
	// 	return err
	// }
	// if ride == nil {
	// 	return RaiseServiceError(errAcceptRideRideNotFound)
	// }
	// if ride.Status != domain.RideStatusAccepted {
	// 	return RaiseServiceError(errStartRideInvalidStatus)
	// }
	// ride.Status = domain.RideStatusInProgres
	// if err := s.rideRepo.SaveRide(ride); err != nil {
	// 	return err
	// }
	return nil
}
