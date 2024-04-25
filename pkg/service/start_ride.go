package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

type StartRideParams struct {
	RideID   string
	DriverID string
}

type StartRide struct {
	rideRepo   repository.RideRepository
	accoutRepo repository.AccountRepository
}

func NewStartRide(rideRepo repository.RideRepository, accountRepo repository.AccountRepository) *StartRide {
	return &StartRide{rideRepo, accountRepo}
}

func (s StartRide) Execute(params StartRideParams) error {
	driverAcc, err := s.accoutRepo.GetAccountByID(params.DriverID)
	if err != nil {
		return err
	}
	if !driverAcc.IsDriver() {
		return errors.New("account is not from a driver")
	}
	ride, err := s.rideRepo.GetRideByID(params.RideID)
	if err != nil {
		return err
	}
	if err := ride.Accept(ride.DriverID()); err != nil {
		return err
	}
	return s.rideRepo.UpdateRide(ride)
}
