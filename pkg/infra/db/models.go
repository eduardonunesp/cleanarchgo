// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type AccountType string

const (
	AccountTypeDriver    AccountType = "driver"
	AccountTypePassenger AccountType = "passenger"
)

func (e *AccountType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccountType(s)
	case string:
		*e = AccountType(s)
	default:
		return fmt.Errorf("unsupported scan type for AccountType: %T", src)
	}
	return nil
}

type NullAccountType struct {
	AccountType AccountType
	Valid       bool // Valid is true if AccountType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccountType) Scan(value interface{}) error {
	if value == nil {
		ns.AccountType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccountType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccountType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AccountType), nil
}

type Account struct {
	ID          pgtype.UUID
	Name        string
	Email       string
	Cpf         string
	CarPlate    pgtype.Text
	AccountType NullAccountType
}

type Position struct {
	PositionID pgtype.UUID
	RideID     pgtype.UUID
	Lat        pgtype.Numeric
	Long       pgtype.Numeric
	Date       pgtype.Timestamp
}

type Ride struct {
	ID          pgtype.UUID
	PassengerID pgtype.UUID
	DriverID    pgtype.UUID
	Status      string
	Fare        pgtype.Numeric
	Distance    pgtype.Numeric
	FromLat     pgtype.Numeric
	FromLong    pgtype.Numeric
	ToLat       pgtype.Numeric
	ToLong      pgtype.Numeric
	Date        pgtype.Timestamp
}
