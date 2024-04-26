package domain

import "github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"

func BuildPosition(opts ...PositionOption) (*Position, error) {
	newPosition := Position{}
	for _, opt := range opts {
		if err := opt(&newPosition); err != nil {
			return nil, err
		}
	}
	return &newPosition, nil
}

func PositionWithNewID() PositionOption {
	return func(p *Position) error {
		p.id = valueobject.MustUUID()
		return nil
	}
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

func PositionWithDateNow() PositionOption {
	return func(p *Position) error {
		p.date = valueobject.DateFromNow()
		return nil
	}
}
