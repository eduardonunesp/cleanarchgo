package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errGetRideErrorNotFound = errors.New("ride not found for given id")
)

type GetRideParams struct {
	RideID string
}

type GetRideResult struct {
	ID             string
	PassengerID    string
	FromLat        string
	FromLong       string
	ToLat          string
	ToLong         string
	Status         string
	PassengerName  string
	PassengerEmail string
}

type GetRide struct {
	rideRepo    repository.RideRepository
	accountRepo repository.AccountRepository
}

func NewGetRide(ridePepo repository.RideRepository, accountRepo repository.AccountRepository) *GetRide {
	return &GetRide{ridePepo, accountRepo}
}

func (g GetRide) Execute(input *GetRideParams) (*GetRideResult, error) {
	ride, err := g.rideRepo.GetRideByID(input.RideID)
	if err != nil {
		return nil, err
	}
	if ride == nil {
		return nil, RaiseServiceError(errGetRideErrorNotFound)
	}
	acc, err := g.accountRepo.GetAccountByID(ride.PassengerID)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, RaiseServiceError(errGetRideErrorNotFound)
	}
	return &GetRideResult{
		ID:             ride.ID,
		PassengerID:    ride.PassengerID,
		FromLat:        ride.FromLat,
		FromLong:       ride.FromLong,
		ToLat:          ride.ToLat,
		ToLong:         ride.ToLong,
		Status:         string(ride.Status),
		PassengerName:  acc.Name,
		PassengerEmail: acc.Email,
	}, nil
}
