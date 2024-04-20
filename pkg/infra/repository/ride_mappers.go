package repository

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/db"
)

func mapDBRideToDomainRide(ride *db.Ride) (*domain.Ride, error) {
	if ride == nil {
		return nil, errors.New("db ride cannot be nil")
	}
	rideStatus, err := domain.BuildRideStatusFromString(ride.Status)
	if err != nil {
		return nil, err
	}
	return domain.BuildRide(
		domain.RideWithIDAndPassengerID(
			fromPgTypeUUIDToString(ride.ID),
			fromPgTypeUUIDToString(ride.PassengerID),
		),
		domain.RideWithFromLatLong(
			fromPgTypeNumericToInt64(ride.FromLat),
			fromPgTypeNumericToInt64(ride.FromLong),
		),
		domain.RideWithFromLatLong(
			fromPgTypeNumericToInt64(ride.ToLat),
			fromPgTypeNumericToInt64(ride.ToLong),
		),
		domain.RideWithStatus(rideStatus),
		domain.RideWithDate(ride.Date.Time),
	)
}

func mapDomainRideToSaveRideParams(ride *domain.Ride) (*db.SaveRideParams, error) {
	if ride == nil {
		return nil, errors.New("domain ride cannot be nil")
	}
	rideUUID, err := mapStringToPgTypeUUID(ride.ID)
	if err != nil {
		return nil, err
	}
	passengerUUID, err := mapStringToPgTypeUUID(ride.PassengerID)
	if err != nil {
		return nil, err
	}
	fromLatNumeric, err := mapInt64ToPgTypeNumeric(ride.FromLat)
	if err != nil {
		return nil, err
	}
	fromLongNumeric, err := mapInt64ToPgTypeNumeric(ride.FromLong)
	if err != nil {
		return nil, err
	}
	toLatNumeric, err := mapInt64ToPgTypeNumeric(ride.ToLat)
	if err != nil {
		return nil, err
	}
	toLongNumeric, err := mapInt64ToPgTypeNumeric(ride.ToLong)
	if err != nil {
		return nil, err
	}
	return &db.SaveRideParams{
		ID:          rideUUID,
		PassengerID: passengerUUID,
		FromLat:     fromLatNumeric,
		FromLong:    fromLongNumeric,
		ToLat:       toLatNumeric,
		ToLong:      toLongNumeric,
		Status:      string(ride.Status),
		Date:        fromTimeToPgTypeTimestamp(ride.Date),
	}, nil
}
