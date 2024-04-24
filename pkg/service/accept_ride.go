package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	ErrAcceptRideIsNotDriver = errors.New("account is not a driver")
)

type AcceptRideRequest struct {
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

func (s AcceptRide) Execute(req AcceptRideRequest) error {
	ride, err := s.rideRepo.GetRideByID(req.RideID)
	if err != nil {
		return err
	}
	account, err := s.accRepo.GetAccountByID(ride.DriverID().String())
	if err != nil {
		return err
	}
	if !account.IsDriver() {
		return RaiseServiceError(ErrAcceptRideIsNotDriver)
	}
	if err := ride.Accept(account.ID()); err != nil {
		return err
	}
	if err := s.rideRepo.UpdateRide(ride); err != nil {
		return err
	}
	return nil
}
