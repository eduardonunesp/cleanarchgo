package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type PositionOption func(position *Position) error

type Position struct {
	PositionID valueobject.UUID
	RideID     valueobject.UUID
	Coord      valueobject.Coord
	Date       valueobject.Date
}

func buildPosition(positionID, rideID, lat, long string, date int64) (*Position, error) {
	var (
		newPosition Position
		err         error
	)
	if newPosition.PositionID, err = valueobject.UUIDFromString(positionID); err != nil {
		newPosition.PositionID = valueobject.MustUUID()
	}
	if newPosition.RideID, err = valueobject.UUIDFromString(rideID); err != nil {
		return nil, err
	}
	if newPosition.Coord, err = valueobject.BuildCoord(lat, long); err != nil {
		return nil, err
	}
	if newPosition.Date, err = valueobject.DateFromUnix(date); err != nil {
		newPosition.Date = valueobject.DateFromNow()
	}
	return &newPosition, nil
}

func CreatePosition(rideID, lat, long string) (*Position, error) {
	newPosition, err := buildPosition("", rideID, lat, long, 0)
	if err != nil {
		return nil, RaiseDomainError(err)
	}
	return newPosition, nil
}

func RestorePosition(positionID, rideID, lat, long string, date int64) (*Position, error) {
	newPosition, err := buildPosition(positionID, rideID, lat, long, date)
	if err != nil {
		return nil, RaiseDomainError(err)
	}
	return newPosition, nil
}
