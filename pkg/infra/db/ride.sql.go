// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: ride.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getRide = `-- name: GetRide :one
SELECT
    id, passenger_id, driver_id, status, fare, distance, from_lat, from_long, to_lat, to_long, date
FROM
    ride
WHERE
    id = $1 LIMIT 1
`

func (q *Queries) GetRide(ctx context.Context, id pgtype.UUID) (Ride, error) {
	row := q.db.QueryRow(ctx, getRide, id)
	var i Ride
	err := row.Scan(
		&i.ID,
		&i.PassengerID,
		&i.DriverID,
		&i.Status,
		&i.Fare,
		&i.Distance,
		&i.FromLat,
		&i.FromLong,
		&i.ToLat,
		&i.ToLong,
		&i.Date,
	)
	return i, err
}

const hasActiveRideByPassengerID = `-- name: HasActiveRideByPassengerID :one
SELECT 
    CASE 
        WHEN status <> 'completed' THEN TRUE
        ELSE FALSE
    END
FROM ride
WHERE 
    passenger_id = $1 
    AND status <> 'completed'
`

func (q *Queries) HasActiveRideByPassengerID(ctx context.Context, passengerID pgtype.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, hasActiveRideByPassengerID, passengerID)
	var column_1 bool
	err := row.Scan(&column_1)
	return column_1, err
}

const saveRide = `-- name: SaveRide :exec
INSERT INTO ride (
    id,
    passenger_id,
    driver_id,
    fare,
    distance,
    from_lat,
    from_long,
    to_lat,
    to_long,
    status,
    date
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
`

type SaveRideParams struct {
	ID          pgtype.UUID
	PassengerID pgtype.UUID
	DriverID    pgtype.UUID
	Fare        pgtype.Numeric
	Distance    pgtype.Numeric
	FromLat     pgtype.Numeric
	FromLong    pgtype.Numeric
	ToLat       pgtype.Numeric
	ToLong      pgtype.Numeric
	Status      string
	Date        pgtype.Timestamp
}

func (q *Queries) SaveRide(ctx context.Context, arg SaveRideParams) error {
	_, err := q.db.Exec(ctx, saveRide,
		arg.ID,
		arg.PassengerID,
		arg.DriverID,
		arg.Fare,
		arg.Distance,
		arg.FromLat,
		arg.FromLong,
		arg.ToLat,
		arg.ToLong,
		arg.Status,
		arg.Date,
	)
	return err
}
