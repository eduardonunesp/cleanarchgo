CREATE TABLE IF NOT EXISTS ride (
	id uuid PRIMARY KEY,
	passenger_id uuid,
	driver_id uuid,
	status TEXT NOT NULL,
	fare NUMERIC,
	distance NUMERIC,
	from_lat NUMERIC,
	from_long NUMERIC,
	to_lat NUMERIC,
	to_long NUMERIC,
	date TIMESTAMP
);