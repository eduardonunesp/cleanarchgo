package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	ErrAcceptRideIsNotDriver = errors.New("account is not a driver")
)

type AcceptRideParams struct {
	RideID   string
	DriverID string
}

type AcceptRide struct {
	rideRepo repository.RideRepository
	accRepo  repository.AccountRepository
}

func NewAcceptRide(rideRepo repository.RideRepository, accRepo repository.AccountRepository) *AcceptRide {
	return &AcceptRide{rideRepo, accRepo}
}

func (s AcceptRide) Execute(input *AcceptRideParams) error {
	ride, err := s.rideRepo.GetRideByID(input.RideID)
	if err != nil {
		return err
	}
	driverAcc, err := s.accRepo.GetAccountByID(ride.DriverID().String())
	if err != nil {
		return err
	}
	if !driverAcc.IsDriver() {
		return RaiseServiceError(ErrAcceptRideIsNotDriver)
	}
	if err := ride.Accept(driverAcc.ID()); err != nil {
		return err
	}
	return s.rideRepo.UpdateRide(ride)
}
