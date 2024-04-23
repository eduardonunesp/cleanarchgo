package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errAcceptRideRideNotFound   = errors.New("ride not found")
	errAcceptRideDriverNotFound = errors.New("driver not found")
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
		return RaiseServiceError(errAcceptRideRideNotFound)
	}
	account, err := s.accRepo.GetAccountByID(ride.DriverID.String())
	if err != nil {
		return err
	}
	if account == nil {
		return RaiseServiceError(errAcceptRideDriverNotFound)
	}
	if err := account.AuthorizeDriver(); err != nil {
		return err
	}
	if err := ride.Accept(account.ID); err != nil {
		return err
	}
	if err := s.rideRepo.SaveRide(ride); err != nil {
		return err
	}
	return nil
}
