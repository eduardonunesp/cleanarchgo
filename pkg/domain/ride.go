package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type rideOption func(*Ride) error

type Ride struct {
	id          valueobject.UUID
	passengerID valueobject.UUID
	driverID    valueobject.UUID
	status      valueobject.RideStatus
	fare        string
	segment     valueobject.Segment
	date        valueobject.Date
}

func (r Ride) ID() valueobject.UUID {
	return r.id
}

func (r Ride) PassengerID() valueobject.UUID {
	return r.passengerID
}

func (r Ride) DriverID() valueobject.UUID {
	return r.driverID
}

func (r Ride) Status() valueobject.RideStatus {
	return r.status
}

func (r Ride) Fare() string {
	return r.fare
}

func (r Ride) Segment() valueobject.Segment {
	return r.segment
}

func (r Ride) Date() valueobject.Date {
	return r.date
}

func (r Ride) IsInProgress() bool {
	return valueobject.RideStatusAs(r.status, valueobject.RideStatusInProgress{})
}

func (r *Ride) Finish(fare string) error {
	newStatus, err := r.status.Finish()
	if err != nil {
		return RaiseDomainError(err)
	}
	r.fare = fare
	r.status = newStatus
	return nil
}

func (r *Ride) Accept(driverID valueobject.UUID) error {
	newStatus, err := r.status.Accept()
	if err != nil {
		return RaiseDomainError(err)
	}
	r.driverID = driverID
	r.status = newStatus
	return nil
}

func RideWithID(id string) rideOption {
	return func(r *Ride) error {
		var err error
		r.id, err = valueobject.UUIDFromString(id)
		return err
	}
}

func RideWithPassengerID(passengerID string) rideOption {
	return func(r *Ride) error {
		var err error
		r.passengerID, err = valueobject.UUIDFromString(passengerID)
		return err
	}
}

func RideWithDriverID(driverID string) rideOption {
	return func(r *Ride) error {
		var err error
		r.driverID, err = valueobject.UUIDFromString(driverID)
		return err
	}
}

func RideWithStatus(status string) rideOption {
	return func(r *Ride) error {
		var err error
		r.status, err = valueobject.BuildRideStatus(status)
		return err
	}
}

func RideWithFare(fare string) rideOption {
	return func(r *Ride) error {
		r.fare = fare
		return nil
	}
}

func RideWithSegment(fromLat, fromLong, toLat, ToLong string) rideOption {
	return func(r *Ride) error {
		var err error
		r.segment, err = valueobject.BuildSegmentFromCoords(fromLat, fromLong, toLat, ToLong)
		return err
	}
}

func RideWithDate(date int64) rideOption {
	return func(r *Ride) error {
		var err error
		r.date, err = valueobject.DateFromUnix(date)
		return err
	}
}

func BuildRide(opts ...rideOption) (*Ride, error) {
	r := &Ride{
		id:     valueobject.MustUUID(),
		date:   valueobject.DateFromNow(),
		status: valueobject.RideStatusRequested{},
	}
	for _, opt := range opts {
		err := opt(r)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}
