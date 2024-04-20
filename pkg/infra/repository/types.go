package repository

import "github.com/eduardonunesp/cleanarchgo/pkg/domain"

//go:generate mockery
type AccountRepositoryRO interface {
	HasAccountByEmail(email string) (bool, error)
	GetAccountByID(id string) (*domain.Account, error)
}

//go:generate mockery
type AccountRepository interface {
	AccountRepositoryRO
	SaveAccount(account *domain.Account) error
}

//go:generate mockery
type RideRepositoryRO interface {
	HasActiveRideByPassengerID(passengerID string) bool
	GetRideByID(rideID string) (*domain.Ride, error)
}

//go:generate mockery
type RideRepository interface {
	RideRepositoryRO
	SaveRide(ride domain.Ride)
}