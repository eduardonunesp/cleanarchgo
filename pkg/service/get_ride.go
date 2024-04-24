package service

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/infra/repository"
)

var (
	errGetRideNotFound = errors.New("ride not found for given id")
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
		return nil, RaiseServiceError(errGetRideNotFound)
	}
	passengerAcc, err := g.accountRepo.GetAccountByID(ride.PassengerID().String())
	if err != nil {
		return nil, err
	}
	if passengerAcc == nil {
		return nil, RaiseServiceError(errGetRideNotFound)
	}
	return &GetRideResult{
		ID:             ride.ID().String(),
		PassengerID:    ride.PassengerID().String(),
		FromLat:        ride.Segment().From().Lat(),
		FromLong:       ride.Segment().From().Long(),
		ToLat:          ride.Segment().To().Lat(),
		ToLong:         ride.Segment().To().Long(),
		Status:         ride.Status().String(),
		PassengerName:  passengerAcc.Name().String(),
		PassengerEmail: passengerAcc.Email().String(),
	}, nil
}
