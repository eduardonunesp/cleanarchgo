package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	errRideEmptyID          = errors.New("ride id cannot be empty")
	errRideEmptyPassengerID = errors.New("passenger id cannot by empty")
	errRideEmptyDriverID    = errors.New("driver id cannot be empty")
)

type RideOption func(ride *Ride) error

type Ride struct {
	ID          string
	PassengerID string
	DriverID    string
	Status      RideStatus
	Fare        string
	Distance    string
	FromLat     string
	FromLong    string
	ToLat       string
	ToLong      string
	Date        time.Time
}

func RideWithID(id string) RideOption {
	return func(ride *Ride) error {
		if id == "" {
			return errRideEmptyID
		}
		ride.ID = id
		return nil
	}
}

func RideWithPassengerID(passengerID string) RideOption {
	return func(ride *Ride) error {
		if passengerID == "" {
			return errRideEmptyPassengerID
		}
		ride.PassengerID = passengerID
		return nil
	}
}

func RideWithDriverID(driverID string) RideOption {
	return func(ride *Ride) error {
		if driverID == "" {
			return errRideEmptyDriverID
		}
		ride.DriverID = driverID
		return nil
	}
}

func RideWithFare(fare string) RideOption {
	return func(ride *Ride) error {
		ride.Fare = fare
		return nil
	}
}

func RideWithDistance(distance string) RideOption {
	return func(ride *Ride) error {
		ride.Distance = distance
		return nil
	}
}

func RideWithFromLatLong(lat, long string) RideOption {
	return func(ride *Ride) error {
		ride.FromLat = lat
		ride.FromLong = long
		return nil
	}
}

func RideWithToLatLong(lat, long string) RideOption {
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

func BuildRide(rideOpts ...RideOption) (*Ride, error) {
	var newRide Ride
	for _, opt := range rideOpts {
		if opt == nil {
			continue
		}
		if err := opt(&newRide); err != nil {
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
