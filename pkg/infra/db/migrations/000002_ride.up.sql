CREATE TYPE status_type AS ENUM ('requested', 'accepted', 'in_progress', 'completed');

CREATE TABLE IF NOT EXISTS ride (
	id uuid PRIMARY KEY,
	passenger_id uuid NOT NULL,
	driver_id uuid,
	status STATUS_TYPE NOT NULL,
	fare NUMERIC,
	distance NUMERIC,
	from_lat NUMERIC NOT NULL,
	from_long NUMERIC NOT NULL,
	to_lat NUMERIC NOT NULL,
	to_long NUMERIC NOT NULL,
	date TIMESTAMP NOT NULL
);