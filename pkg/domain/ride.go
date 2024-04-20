package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrRideEmptyID          = errors.New("ride id cannot be empty")
	ErrRideEmptyPassengerID = errors.New("passenger id cannot by empty")
)

type RideOption func(ride *Ride) error

type Ride struct {
	ID          string
	PassengerID string
	DriverID    string
	Status      RideStatus
	Fare        int64
	FromLat     int64
	FromLong    int64
	ToLat       int64
	ToLong      int64
	Date        time.Time
}

func RideWithIDAndPassengerID(rideID, passengerID string) RideOption {
	return func(ride *Ride) error {
		if rideID == "" {
			return ErrRideEmptyID
		}
		if passengerID == "" {
			return ErrRideEmptyPassengerID
		}
		ride.ID = rideID
		ride.PassengerID = passengerID
		return nil
	}
}

func RideWithPassengerID(passengerID string) RideOption {
	return func(ride *Ride) error {
		if passengerID == "" {
			return ErrRideEmptyPassengerID
		}
		ride.PassengerID = passengerID
		return nil
	}
}

func RideWithFromLatLong(lat, long int64) RideOption {
	return func(ride *Ride) error {
		ride.FromLat = lat
		ride.FromLong = long
		return nil
	}
}

func RideWithToLatLong(lat, long int64) RideOption {
	return func(ride *Ride) error {
		ride.ToLat = lat
		ride.ToLong = long
		return nil
	}
}

func RideWithStatus(rideStatus RideStatus) RideOption {
	return func(ride *Ride) error {
		ride.Status = rideStatus
		return nil
	}
}

func RideWithDate(timeNow time.Time) RideOption {
	return func(ride *Ride) error {
		ride.Date = timeNow
		return nil
	}
}

func BuildRide(accOpts ...RideOption) (*Ride, error) {
	var newRide Ride
	for _, accOpt := range accOpts {
		if accOpt == nil {
			continue
		}
		if err := accOpt(&newRide); err != nil {
			return nil, RaiseDomainError(fmt.Errorf("failed to build account: %w", err))
		}
	}
	rideApplyDefaultParams(&newRide)
	return &newRide, nil
}

func rideApplyDefaultParams(newRide *Ride) {
	if newRide.ID == "" {
		newRide.ID = uuid.Must(uuid.NewRandom()).String()
	}
	if newRide.Status == "" {
		newRide.Status = RideStatusRequested
	}
	if newRide.Date.IsZero() {
		newRide.Date = time.Now()
	}
}
