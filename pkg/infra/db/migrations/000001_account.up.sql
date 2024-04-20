CREATE TABLE IF NOT EXISTS account (
	id uuid PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	cpf TEXT NOT NULL,
	car_plate TEXT NULL,
	is_passenger BOOLEAN NOT NULL DEFAULT false,
	is_driver BOOLEAN NOT NULL DEFAULT false
);