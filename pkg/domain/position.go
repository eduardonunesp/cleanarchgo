package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type PositionOption func(position *Position) error

type Position struct {
	id     valueobject.UUID
	rideID valueobject.UUID
	coord  valueobject.Coord
	date   valueobject.Date
}

func CreatePosition(rideID, lat, long string) (*Position, error) {
	var (
		newPosition Position
		err         error
	)
	newPosition.id = valueobject.MustUUID()
	if newPosition.rideID, err = valueobject.UUIDFromString(rideID); err != nil {
		return nil, err
	}
	if newPosition.coord, err = valueobject.BuildCoord(lat, long); err != nil {
		return nil, err
	}
	newPosition.date = valueobject.DateFromNow()
	return &newPosition, nil
}

func RestorePosition(positionID, rideID, lat, long string, date int64) (*Position, error) {
	var (
		newPosition Position
		err         error
	)
	if newPosition.id, err = valueobject.UUIDFromString(positionID); err != nil {
		return nil, err
	}
	if newPosition.rideID, err = valueobject.UUIDFromString(rideID); err != nil {
		return nil, err
	}
	if newPosition.coord, err = valueobject.BuildCoord(lat, long); err != nil {
		return nil, err
	}
	if newPosition.date, err = valueobject.DateFromUnix(date); err != nil {
		return nil, err
	}
	return &newPosition, nil
}

func (p Position) ID() valueobject.UUID {
	return p.id
}

func (p Position) RideID() valueobject.UUID {
	return p.rideID
}

func (p Position) Coord() valueobject.Coord {
	return p.coord
}

func (p Position) Date() valueobject.Date {
	return p.date
}
