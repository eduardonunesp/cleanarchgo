package domain

import (
	"fmt"

	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type PositionOption func(position *Position) error

type Position struct {
	PositionID valueobject.UUID
	RideID     valueobject.UUID
	Coord      valueobject.Coord
	Date       valueobject.Date
}

func WithPositionID(positionID string) PositionOption {
	return func(position *Position) error {
		var err error
		if position.PositionID, err = valueobject.UUIDFromString(positionID); err != nil {
			return err
		}
		return nil
	}
}

func WithRideID(rideID string) PositionOption {
	return func(position *Position) error {
		var err error
		if position.RideID, err = valueobject.UUIDFromString(rideID); err != nil {
			return err
		}
		return nil
	}
}

func WithLatLong(lat, long string) PositionOption {
	return func(position *Position) error {
		var err error
		if position.Coord, err = valueobject.NewCoord(lat, long); err != nil {
			return err
		}
		return nil
	}
}

func WithDate(date int64) PositionOption {
	return func(position *Position) error {
		position.Date = valueobject.DateFromInt64(date)
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
			return nil, RaiseDomainError(fmt.Errorf("failed to build Position: %w", err))
		}
	}
	positionApplyDefaultParams(&newPosition)
	return &newPosition, nil
}

func positionApplyDefaultParams(newPosition *Position) {
	if newPosition.PositionID.String() == "" {
		newPosition.PositionID = valueobject.MustUUID()
	}
	if newPosition.Date.IsZero() {
		newPosition.Date = valueobject.DateFromNow()
	}
}
