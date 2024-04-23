package domain

import (
	"fmt"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type RideOption func(ride *Ride) error

type Ride struct {
	ID          valueobject.UUID
	PassengerID valueobject.UUID
	DriverID    valueobject.UUID
	Status      valueobject.RideStatus
	Fare        string
	FromCoord   valueobject.Coord
	ToCoord     valueobject.Coord
	Date        valueobject.Date
}

func RideWithID(id string) RideOption {
	return func(ride *Ride) error {
		var err error
		if ride.ID, err = valueobject.UUIDFromString(id); err != nil {
			return err
		}
		return nil
	}
}

func RideWithPassengerID(passengerID string) RideOption {
	return func(ride *Ride) error {
		var err error
		if ride.PassengerID, err = valueobject.UUIDFromString(passengerID); err != nil {
			return err
		}
		return nil
	}
}

func RideWithDriverID(driverID string) RideOption {
	return func(ride *Ride) error {
		var err error
		if ride.DriverID, err = valueobject.UUIDFromString(driverID); err != nil {
			return err
		}
		return nil
	}
}

func RideWithFare(fare string) RideOption {
	return func(ride *Ride) error {
		ride.Fare = fare
		return nil
	}
}

func RideWithFromLatLong(lat, long string) RideOption {
	return func(ride *Ride) error {
		var err error
		if ride.FromCoord, err = valueobject.NewCoord(lat, long); err != nil {
			return err
		}
		return nil
	}
}

func RideWithToLatLong(lat, long string) RideOption {
	return func(ride *Ride) error {
		var err error
		if ride.ToCoord, err = valueobject.NewCoord(lat, long); err != nil {
			return err
		}
		return nil
	}
}

func RideWithStatus(rideStatus string) RideOption {
	return func(ride *Ride) error {
		var err error
		if ride.Status, err = valueobject.BuildRideStatus(rideStatus); err != nil {
			return err
		}
		return nil
	}
}

func RideWithDate(timeNow int64) RideOption {
	return func(ride *Ride) error {
		ride.Date = valueobject.DateFromInt64(timeNow)
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
	if newRide.ID.String() == "" {
		newRide.ID = valueobject.MustUUID()
	}
	if newRide.Date.IsZero() {
		newRide.Date = valueobject.DateFromNow()
	}
}

func (r *Ride) Accept(driverID valueobject.UUID) error {
	newStatus, err := r.Status.Accept()
	if err != nil {
		return RaiseDomainError(err)
	}
	r.DriverID = driverID
	r.Status = newStatus
	return nil
}
