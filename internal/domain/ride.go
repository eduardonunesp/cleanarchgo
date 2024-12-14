package domain

import (
	"github.com/eduardonunesp/cleanarchgo/internal/domain/valueobject"
)

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
	return r.status.Is(valueobject.StatusInProgress)
}

func (r Ride) Finish(fare string) (Ride, error) {
	newStatus, err := r.status.Transition(valueobject.ActionFinish)
	if err != nil {
		return r, RaiseDomainError(err)
	}
	r.fare = fare
	r.status = newStatus
	return r, nil
}

func (r Ride) Accept(driverID valueobject.UUID) (Ride, error) {
	newStatus, err := r.status.Transition(valueobject.ActionAccept)
	if err != nil {
		return r, RaiseDomainError(err)
	}
	r.driverID = driverID
	r.status = newStatus
	return r, nil
}
