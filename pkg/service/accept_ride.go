package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errRideNotFound         = errors.New("ride not found")
	errRideNotRequested     = errors.New("ride not in requested status")
	errAccountNotIsDriver   = errors.New("account isn't a driver")
	errDriverIsNotAvailable = errors.New("driver is not available")
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
	if ride == nil {
		return RaiseServiceError(errRideNotFound)
	}
	driverAcc, err := s.accRepo.GetAccountByID(ride.DriverID)
	if err != nil {
		return err
	}
	if driverAcc == nil {
		return RaiseServiceError(errRideNotFound)
	}
	if driverAcc.IsDriver == false {
		return RaiseServiceError(errAccountNotIsDriver)
	}
	if ride.Status != domain.RideStatusRequested {
		return RaiseServiceError(errRideNotRequested)
	}
	driverFree, err := s.accRepo.IsDriverFreeByDriverID(req.DriverID)
	if err != nil {
		return err
	}
	if !driverFree {
		return RaiseServiceError(errDriverIsNotAvailable)
	}
	ride.DriverID = req.DriverID
	ride.Status = domain.RideStatusAccepted
	if err := s.rideRepo.SaveRide(ride); err != nil {
		return err
	}
	return nil
}
