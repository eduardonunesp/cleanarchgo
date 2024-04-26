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
