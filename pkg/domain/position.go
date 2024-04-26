package domain

import (
	"github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"
)

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
