-- name: GetRide :one
SELECT
    *
FROM
    ride
WHERE
    id = $1 LIMIT 1;

-- name: HasActiveRideByPassengerID :one
SELECT
	CASE WHEN status <> 'completed' THEN
		TRUE
	ELSE
		FALSE
	END
FROM
	ride
WHERE
	passenger_id = $1
ORDER BY
	date DESC
LIMIT 1;

-- name: HasActiveRideByDriverID :one
SELECT
	CASE WHEN status IN('accepted', 'in_progress') THEN
		TRUE
	ELSE
		FALSE
	END
FROM
	ride
WHERE
	driver_id = $1
ORDER BY
	date DESC
LIMIT 1;

-- name: SaveRide :exec
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
);