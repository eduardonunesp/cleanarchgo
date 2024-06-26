-- name: GetAccount :one
SELECT
	*
FROM
	account
WHERE
	id = $1
LIMIT 1;

-- name: HasAccountByEmail :one
SELECT
	CASE
		WHEN count(id) > 0 THEN TRUE
		ELSE FALSE
	END
FROM account
WHERE
	email = $1
LIMIT 1;

-- name: SaveAccount :exec
INSERT INTO account (
	id,
	name,
	email,
	cpf,
	car_plate,
	account_type
) VALUES (
	$1, $2, $3, $4, $5, $6
);
