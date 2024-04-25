CREATE TYPE status_type AS ENUM ('requested', 'accepted', 'in_progress', 'completed');

CREATE TABLE IF NOT EXISTS ride (
	id uuid PRIMARY KEY,
	passenger_id uuid,
	driver_id uuid,
	status STATUS_TYPE,
	fare NUMERIC,
	distance NUMERIC,
	from_lat NUMERIC,
	from_long NUMERIC,
	to_lat NUMERIC,
	to_long NUMERIC,
	date TIMESTAMP
);