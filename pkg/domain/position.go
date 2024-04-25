package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

type PositionOption func(*Position) error

type Position struct {
	id     valueobject.UUID
	rideID valueobject.UUID
	coord  valueobject.Coord
	date   valueobject.Date
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

func PositionWithID(id string) PositionOption {
	return func(p *Position) error {
		var err error
		p.id, err = valueobject.UUIDFromString(id)
		return err
	}
}

func PositionWithRideID(rideID string) PositionOption {
	return func(p *Position) error {
		var err error
		p.rideID, err = valueobject.UUIDFromString(rideID)
		return err
	}
}

func PositionWithCoord(lat, long string) PositionOption {
	return func(p *Position) error {
		var err error
		p.coord, err = valueobject.BuildCoord(lat, long)
		return err
	}
}

func PositionWithDate(date int64) PositionOption {
	return func(p *Position) error {
		var err error
		p.date, err = valueobject.DateFromUnix(date)
		return err
	}
}

func BuildPosition(opts ...PositionOption) (*Position, error) {
	newPosition := Position{
		id:   valueobject.MustUUID(),
		date: valueobject.DateFromNow(),
	}
	for _, opt := range opts {
		if err := opt(&newPosition); err != nil {
			return nil, err
		}
	}
	return &newPosition, nil
}
