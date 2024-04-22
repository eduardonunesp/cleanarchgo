package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errAcceptRideNotFound           = errors.New("ride not found")
	errAcceptRideNotRequested       = errors.New("ride not in requested status")
	errAcceptRideNotDriver          = errors.New("account isn't a driver")
	errAcceptRideDriverNotAvailable = errors.New("driver is not available")
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
		return RaiseServiceError(errAcceptRideNotFound)
	}
	driverAcc, err := s.accRepo.GetAccountByID(ride.DriverID)
	if err != nil {
		return err
	}
	if driverAcc == nil {
		return RaiseServiceError(errAcceptRideNotFound)
	}
	if !driverAcc.IsDriver {
		return RaiseServiceError(errAcceptRideNotDriver)
	}
	if ride.Status != domain.RideStatusRequested {
		return RaiseServiceError(errAcceptRideNotRequested)
	}
	driverFree, err := s.rideRepo.HasActiveRideByDriverID(req.DriverID)
	if err != nil {
		return err
	}
	if !driverFree {
		return RaiseServiceError(errAcceptRideDriverNotAvailable)
	}
	ride.DriverID = req.DriverID
	ride.Status = domain.RideStatusAccepted
	if err := s.rideRepo.SaveRide(ride); err != nil {
		return err
	}
	return nil
}
