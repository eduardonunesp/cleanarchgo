package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
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

func CreateRide(
	passengerID, fare,
	fromLat, fromLong, toLat, toLong string,
) (*Ride, error) {
	var (
		newRide Ride
		err     error
	)
	newRide.id = valueobject.MustUUID()
	if newRide.passengerID, err = valueobject.UUIDFromString(passengerID); err != nil {
		return nil, err
	}
	newRide.status, _ = valueobject.BuildRideStatus(valueobject.RideStatusRequested{}.String())
	newRide.fare = fare
	if newRide.segment, err = valueobject.BuildSegmentFromCoords(
		fromLat, fromLong, toLat, toLong,
	); err != nil {
		return nil, err
	}
	newRide.date = valueobject.DateFromNow()
	return &newRide, nil
}

func RestoreRide(
	id, passengerID, driverID, fare,
	fromLat, fromLong, toLat, toLong, status string,
	date int64,
) (*Ride, error) {
	var (
		newRide Ride
		err     error
	)
	if newRide.id, err = valueobject.UUIDFromString(id); err != nil {
		return nil, err
	}
	if newRide.passengerID, err = valueobject.UUIDFromString(passengerID); err != nil {
		return nil, err
	}
	if newRide.driverID, err = valueobject.UUIDFromString(driverID); err != nil {
		return nil, err
	}
	if newRide.status, err = valueobject.BuildRideStatus(status); err != nil {
		return nil, err
	}
	newRide.fare = fare
	if newRide.segment, err = valueobject.BuildSegmentFromCoords(
		fromLat, fromLong, toLat, toLong,
	); err != nil {
		return nil, err
	}
	if newRide.date, err = valueobject.DateFromUnix(date); err != nil {
		return nil, err
	}
	return &newRide, nil
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
