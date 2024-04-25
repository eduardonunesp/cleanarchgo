package repository

import (
	"errors"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/db"
	"github.com/jackc/pgx/v5/pgtype"
)

func mapDBRideToDomainRide(ride *db.Ride) (*domain.Ride, error) {
	if ride == nil {
		return nil, errors.New("db ride cannot be nil")
	}
	return domain.BuildRide(
		domain.RideWithID(fromPgTypeUUIDToString(ride.ID)),
		domain.RideWithPassengerID(fromPgTypeUUIDToString(ride.PassengerID)),
		domain.RideWithDriverID(fromPgTypeUUIDToString(ride.DriverID)),
		domain.RideWithFare(fromPgTypeNumericToString(ride.Fare)),
		domain.RideWithSegment(
			fromPgTypeNumericToString(ride.FromLat),
			fromPgTypeNumericToString(ride.FromLong),
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
	rideUUID, err := mapStringToPgTypeUUID(ride.ID().String())
	if err != nil {
		return nil, err
	}
	passengerUUID, err := mapStringToPgTypeUUID(ride.PassengerID().String())
	if err != nil {
		return nil, err
	}
	var driverUUID pgtype.UUID
	if ride.DriverID().String() != "" {
		driverUUID, err = mapStringToPgTypeUUID(ride.DriverID().String())
		if err != nil {
			return nil, err
		}
	}
	fare, err := mapStringToPgTypeNumeric(ride.Fare())
	if err != nil {
		return nil, err
	}
	fromLatNumeric, err := mapStringToPgTypeNumeric(ride.Segment().From().Lat())
	if err != nil {
		return nil, err
	}
	fromLongNumeric, err := mapStringToPgTypeNumeric(ride.Segment().From().Long())
	if err != nil {
		return nil, err
	}
	toLatNumeric, err := mapStringToPgTypeNumeric(ride.Segment().To().Lat())
	if err != nil {
		return nil, err
	}
	toLongNumeric, err := mapStringToPgTypeNumeric(ride.Segment().To().Long())
	if err != nil {
		return nil, err
	}
	return &db.SaveRideParams{
		ID:          rideUUID,
		DriverID:    driverUUID,
		Fare:        fare,
		PassengerID: passengerUUID,
		FromLat:     fromLatNumeric,
		FromLong:    fromLongNumeric,
		ToLat:       toLatNumeric,
		ToLong:      toLongNumeric,
		Status:      ride.Status().String(),
		Date:        fromTimeToPgTypeTimestamp(ride.Date().Unix()),
	}, nil
}

func mapDomainRideToUpdateRideParams(ride *domain.Ride) (*db.UpdateRideParams, error) {
	if ride == nil {
		return nil, errors.New("domain ride cannot be nil")
	}
	rideUUID, err := mapStringToPgTypeUUID(ride.ID().String())
	if err != nil {
		return nil, err
	}
	driverUUID, err := mapStringToPgTypeUUID(ride.DriverID().String())
	if err != nil {
		return nil, err
	}
	return &db.UpdateRideParams{
		ID:       rideUUID,
		DriverID: driverUUID,
		Status:   ride.Status().String(),
	}, nil
}
