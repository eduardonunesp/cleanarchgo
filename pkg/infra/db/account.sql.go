// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: account.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAccount = `-- name: GetAccount :one
SELECT 
    id, name, email, cpf, car_plate, is_passenger, is_driver
FROM 
    account
WHERE 
    id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id pgtype.UUID) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Cpf,
		&i.CarPlate,
		&i.IsPassenger,
		&i.IsDriver,
	)
	return i, err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT
    email
FROM 
    account
WHERE 
    email = $1
LIMIT 1
`

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (string, error) {
	row := q.db.QueryRow(ctx, getAccountByEmail, email)
	err := row.Scan(&email)
	return email, err
}

const hasAccountByEmail = `-- name: HasAccountByEmail :one
SELECT 
    CASE 
        WHEN count(id) > 0 THEN TRUE
        ELSE FALSE
    END
FROM account
WHERE 
    email = $1
`

func (q *Queries) HasAccountByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, hasAccountByEmail, email)
	var column_1 bool
	err := row.Scan(&column_1)
	return column_1, err
}

const saveAccount = `-- name: SaveAccount :exec
INSERT INTO account (
    id,
    name,
    email,
    cpf,
    car_plate,
    is_passenger,
    is_driver
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
`

type SaveAccountParams struct {
	ID          pgtype.UUID
	Name        string
	Email       string
	Cpf         string
	CarPlate    pgtype.Text
	IsPassenger bool
	IsDriver    bool
}

func (q *Queries) SaveAccount(ctx context.Context, arg SaveAccountParams) error {
	_, err := q.db.Exec(ctx, saveAccount,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Cpf,
		arg.CarPlate,
		arg.IsPassenger,
		arg.IsDriver,
	)
	return err
}
