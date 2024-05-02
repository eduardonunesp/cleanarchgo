package domain

import "github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"

type PositionCreateParams struct {
	RideID string
	CoordX string
	CoordY string
}

type PositionRestoreParams struct {
	ID     string
	RideID string
	CoordX string
	CoordY string
	Date   int64
}

func PositionCreate(params PositionCreateParams) (*Position, error) {
	var (
		newPosition Position
		err         error
	)
	newPosition.id = valueobject.MustUUID()
	if newPosition.rideID, err = valueobject.UUIDFromString(params.RideID); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newPosition.coord, err = valueobject.BuildCoord(params.CoordX, params.CoordY); err != nil {
		return nil, RaiseDomainError(err)
	}
	newPosition.date = valueobject.DateFromNow()
	return &newPosition, err
}

func PositionRestore(params PositionRestoreParams) (*Position, error) {
	var (
		newPosition Position
		err         error
	)
	if newPosition.id, err = valueobject.UUIDFromString(params.ID); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newPosition.rideID, err = valueobject.UUIDFromString(params.RideID); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newPosition.coord, err = valueobject.BuildCoord(params.CoordX, params.CoordY); err != nil {
		return nil, RaiseDomainError(err)
	}
	if newPosition.date, err = valueobject.DateFromUnix(params.Date); err != nil {
		return nil, RaiseDomainError(err)
	}
	return &newPosition, err
}
