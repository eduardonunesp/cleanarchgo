package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PositionOption func(position *Position) error

type Position struct {
	PositionID string
	RideID     string
	Lat        string
	Long       string
	Date       time.Time
}

func WithPositionID(positionID string) PositionOption {
	return func(position *Position) error {
		position.PositionID = positionID
		return nil
	}
}

func WithRideID(rideID string) PositionOption {
	return func(position *Position) error {
		position.RideID = rideID
		return nil
	}
}

func WithLat(lat string) PositionOption {
	return func(position *Position) error {
		position.Lat = lat
		return nil
	}
}

func WithLong(long string) PositionOption {
	return func(position *Position) error {
		position.Long = long
		return nil
	}
}

func WithDate(date time.Time) PositionOption {
	return func(position *Position) error {
		position.Date = date
		return nil
	}
}

func BuildPosition(posOpts ...PositionOption) (*Position, error) {
	var newPosition Position
	for _, opt := range posOpts {
		if opt == nil {
			continue
		}
		if err := opt(&newPosition); err != nil {
			return nil, RaiseDomainError(fmt.Errorf("failed to build account: %w", err))
		}
	}
	positionApplyDefaultParams(&newPosition)
	return &newPosition, nil
}

func positionApplyDefaultParams(newPosition *Position) {
	if newPosition.PositionID == "" {
		newPosition.PositionID = uuid.Must(uuid.NewRandom()).String()
	}
	if newPosition.Date.IsZero() {
		newPosition.Date = time.Now()
	}
}
