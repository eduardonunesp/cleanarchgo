package domain

import (
	"github.com/eduardonunesp/cleanarchgo/internal/domain/valueobject"
)

type RideCreateParams struct {
	PassengerID string
	InitialX    string
	InitialY    string
	FinalX      string
	FinalY      string
}

type RideRestoreParams struct {
	ID          string
	PassengerID string
	DriverID    string
	Status      string
	InitialX    string
	InitialY    string
	FinalX      string
	FinalY      string
	Date        int64
	Fare        string
}

func RideCreate(params RideCreateParams) (*Ride, error) {
	var (
		newRide Ride
		err     error
	)
	newRide.id = valueobject.MustUUID()
	if newRide.passengerID, err = valueobject.UUIDFromString(params.PassengerID); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid passenger id")
	}
	if newRide.segment, err = valueobject.BuildSegmentFromCoordinates(
		params.InitialX,
		params.InitialY,
		params.FinalX,
		params.FinalY,
	); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid lat/long coordinates")
	}
	newRide.status, _ = valueobject.BuildRideStatus(valueobject.StatusRequested)
	newRide.date = valueobject.DateFromNow()
	return &newRide, err
}

func RideRestore(params RideRestoreParams) (*Ride, error) {
	var (
		newRide Ride
		err     error
	)
	if newRide.id, err = valueobject.UUIDFromString(params.ID); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid ride id")
	}
	if newRide.passengerID, err = valueobject.UUIDFromString(params.PassengerID); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid passenger id")
	}
	if newRide.driverID, err = valueobject.UUIDFromString(params.DriverID); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid driver id")
	}
	if newRide.segment, err = valueobject.BuildSegmentFromCoordinates(
		params.InitialX,
		params.InitialY,
		params.FinalX,
		params.FinalY,
	); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid lat long coords")
	}
	if newRide.status, err = valueobject.BuildRideStatus(valueobject.Status(params.Status)); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid ride status")
	}
	if newRide.date, err = valueobject.DateFromUnix(params.Date); err != nil {
		return nil, RaiseDomainErrorWith(err, "invalid ride date")
	}
	newRide.fare = params.Fare
	return &newRide, err
}
