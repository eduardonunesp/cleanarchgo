package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errRequestRideNotFound     = errors.New("passenger not found")
	errRequestRideNotPassenger = errors.New("account is not a passenger")
	errRequestRideActiveFound  = errors.New("passenger has an active ride")
)

type RequestRideParams struct {
	PassengerID string
	FromLat     string
	FromLong    string
	ToLat       string
	ToLong      string
}

type RequestRideResult struct {
	RideID string
}

type RequestRide struct {
	rideRepo repository.RideRepository
	accRepo  repository.AccountRepository
}

func NewRequestRide(rideRepo repository.RideRepository, accRepo repository.AccountRepository) *RequestRide {
	return &RequestRide{rideRepo, accRepo}
}

func (r *RequestRide) Execute(params *RequestRideParams) (*RequestRideResult, error) {
	account, err := r.accRepo.GetAccountByID(params.PassengerID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, RaiseServiceError(errRequestRideNotFound)
	}
	if !account.IsPassenger {
		return nil, RaiseServiceError(errRequestRideNotPassenger)
	}
	hasActiveRide, err := r.rideRepo.HasActiveRideByPassengerID(params.PassengerID)
	if err != nil {
		return nil, err
	}
	if hasActiveRide {
		return nil, RaiseServiceError(errRequestRideActiveFound)
	}
	ride, err := domain.BuildRide(
		domain.RideWithPassengerID(params.PassengerID),
		domain.RideWithFromLatLong(params.FromLat, params.FromLong),
		domain.RideWithToLatLong(params.ToLat, params.ToLong),
	)
	if err != nil {
		return nil, err
	}
	if err := r.rideRepo.SaveRide(ride); err != nil {
		return nil, err
	}
	return &RequestRideResult{
		RideID: ride.ID.String(),
	}, nil
}
