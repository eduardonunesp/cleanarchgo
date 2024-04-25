CREATE TYPE account_type AS ENUM ('driver', 'passenger');

CREATE TABLE IF NOT EXISTS account (
	id uuid PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	cpf TEXT NOT NULL,
	car_plate TEXT NULL,
	account_type ACCOUNT_TYPE
);