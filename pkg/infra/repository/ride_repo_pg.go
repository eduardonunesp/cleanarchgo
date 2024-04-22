package repository

import (
	"context"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain"
	"github.com/eduardonunesp/cleanarchgo/pkg/infra/db"
	"github.com/jackc/pgx/v5"
)

type RideRepositoryPG struct {
	conn *pgx.Conn
}

func NewRideRepositoryPG(connStr string) *RideRepositoryPG {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	return &RideRepositoryPG{conn}
}

func (r RideRepositoryPG) HasActiveRideByPassengerID(passengerID string) (bool, error) {
	passengerUUID, err := mapStringToPgTypeUUID(passengerID)
	if err != nil {
		return false, RaiseRepositoryError(err)
	}
	queries := db.New(r.conn)
	result, err := queries.HasActiveRideByPassengerID(context.Background(), passengerUUID)
	if err != nil {
		return false, RaiseRepositoryError(err)
	}
	return result, nil
}

func (r RideRepositoryPG) HasActiveRideByDriverID(driverID string) (bool, error) {
	driverUUID, err := mapStringToPgTypeUUID(driverID)
	if err != nil {
		return false, RaiseRepositoryError(err)
	}
	queries := db.New(r.conn)
	result, err := queries.HasActiveRideByDriverID(context.Background(), driverUUID)
	if err != nil {
		return false, RaiseRepositoryError(err)
	}
	return result, nil
}

func (r RideRepositoryPG) GetRideByID(rideID string) (*domain.Ride, error) {
	rideUUID, err := mapStringToPgTypeUUID(rideID)
	if err != nil {
		return nil, RaiseRepositoryError(err)
	}
	queries := db.New(r.conn)
	result, err := queries.GetRide(context.Background(), rideUUID)
	if err != nil {
		return nil, RaiseRepositoryError(err)
	}
	return mapDBRideToDomainRide(&result)
}

func (r RideRepositoryPG) SaveRide(ride *domain.Ride) error {
	saveRideParams, err := mapDomainRideToSaveRideParams(ride)
	if err != nil {
		return RaiseRepositoryError(err)
	}
	queries := db.New(r.conn)
	err = queries.SaveRide(context.Background(), *saveRideParams)
	if err != nil {
		return RaiseRepositoryError(err)
	}
	return nil
}
