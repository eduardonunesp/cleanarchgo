CREATE TABLE IF NOT EXISTS position (
    position_id UUID,
    ride_id UUID,
    lat NUMERIC,
    long NUMERIC,
    date TIMESTAMP
);