package repository

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/db"
)

func mapDBRideToDomainRide(ride *db.Ride) (*domain.Ride, error) {
	return domain.BuildRide(
		domain.RideWithID(
			fromPgTypeUUIDToString(ride.ID),
		),
		domain.RideWithPassengerID(
			fromPgTypeUUIDToString(ride.PassengerID),
		),
		domain.RideWithDriverID(
			fromPgTypeUUIDToString(ride.DriverID),
		),
		domain.RideWithFare(
			fromPgTypeNumericToString(ride.Fare),
		),
		domain.RideWithFromLatLong(
			fromPgTypeNumericToString(ride.FromLat),
			fromPgTypeNumericToString(ride.FromLong),
		),
		domain.RideWithToLatLong(
			fromPgTypeNumericToString(ride.ToLat),
			fromPgTypeNumericToString(ride.ToLong),
		),
		domain.RideWithStatus(ride.Status),
		domain.RideWithDate(ride.Date.Time.Unix()),
	)
}

func mapDomainRideToSaveRideParams(ride *domain.Ride) (*db.SaveRideParams, error) {
	if ride == nil {
		return nil, errors.New("domain ride cannot be nil")
	}
	rideUUID, err := mapStringToPgTypeUUID(ride.ID.String())
	if err != nil {
		return nil, err
	}
	passengerUUID, err := mapStringToPgTypeUUID(ride.PassengerID.String())
	if err != nil {
		return nil, err
	}
	driverUUID, err := mapStringToPgTypeUUID(ride.DriverID.String())
	if err != nil {
		return nil, err
	}
	fare, err := mapStringToPgTypeNumeric(ride.Fare)
	if err != nil {
		return nil, err
	}
	// distance, err := mapStringToPgTypeNumeric(ride.Distance)
	// if err != nil {
	// 	return nil, err
	// }
	fromLatNumeric, err := mapStringToPgTypeNumeric(ride.FromCoord.Lat())
	if err != nil {
		return nil, err
	}
	fromLongNumeric, err := mapStringToPgTypeNumeric(ride.FromCoord.Lat())
	if err != nil {
		return nil, err
	}
	toLatNumeric, err := mapStringToPgTypeNumeric(ride.ToCoord.Lat())
	if err != nil {
		return nil, err
	}
	toLongNumeric, err := mapStringToPgTypeNumeric(ride.ToCoord.Long())
	if err != nil {
		return nil, err
	}
	return &db.SaveRideParams{
		ID:       rideUUID,
		DriverID: driverUUID,
		Fare:     fare,
		// Distance:    distance,
		PassengerID: passengerUUID,
		FromLat:     fromLatNumeric,
		FromLong:    fromLongNumeric,
		ToLat:       toLatNumeric,
		ToLong:      toLongNumeric,
		Status:      ride.Status.String(),
		Date:        fromTimeToPgTypeTimestamp(ride.Date.Int64()),
	}, nil
}
