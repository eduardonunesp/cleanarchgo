package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type Ride struct {
	ID          valueobject.UUID
	PassengerID valueobject.UUID
	DriverID    valueobject.UUID
	Status      valueobject.RideStatus
	Fare        string
	Segment     valueobject.Segment
	Date        valueobject.Date
}

func CreateRide(
	passengerID, fare,
	fromLat, fromLong, toLat, toLong string,
) (*Ride, error) {
	var (
		newRide Ride
		err     error
	)
	newRide.ID = valueobject.MustUUID()
	if newRide.PassengerID, err = valueobject.UUIDFromString(passengerID); err != nil {
		return nil, err
	}
	newRide.Status, _ = valueobject.BuildRideStatus(valueobject.RideStatusRequested{}.String())
	newRide.Fare = fare
	if newRide.Segment, err = valueobject.BuildSegmentFromCoords(
		fromLat, fromLong, toLat, toLong,
	); err != nil {
		return nil, err
	}
	newRide.Date = valueobject.DateFromNow()
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
	if newRide.ID, err = valueobject.UUIDFromString(id); err != nil {
		return nil, err
	}
	if newRide.PassengerID, err = valueobject.UUIDFromString(passengerID); err != nil {
		return nil, err
	}
	if newRide.DriverID, err = valueobject.UUIDFromString(driverID); err != nil {
		return nil, err
	}
	if newRide.Status, err = valueobject.BuildRideStatus(status); err != nil {
		return nil, err
	}
	newRide.Fare = fare
	if newRide.Segment, err = valueobject.BuildSegmentFromCoords(
		fromLat, fromLong, toLat, toLong,
	); err != nil {
		return nil, err
	}
	if newRide.Date, err = valueobject.DateFromUnix(date); err != nil {
		return nil, err
	}
	return &newRide, nil
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
