package valueobject

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidLatCoord   = errors.New("invalid latitude coord")
	ErrInvalidLongCoord  = errors.New("invalid longitude coord")
	ErrFailedToParseLat  = errors.New("failed to parse latitude")
	ErrFailedToParseLong = errors.New("failed to parse longitude")
)

const (
	minLat  = -90
	maxLat  = 90
	minLong = -180
	maxLong = 180
)

type Coord struct {
	x string
	y string
}

func BuildCoord(lat, long string) (Coord, error) {
	ilat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return Coord{}, ErrFailedToParseLat
	}
	ilong, err := strconv.ParseFloat(long, 64)
	if err != nil {
		return Coord{}, ErrFailedToParseLong
	}
	if ilat < minLat || ilat > maxLat {
		return Coord{}, ErrInvalidLatCoord
	}
	if ilong < minLong || ilong > maxLong {
		return Coord{}, ErrInvalidLongCoord
	}
	return Coord{lat, long}, nil
}

func (c Coord) Lat() string {
	return c.x
}

func (c Coord) Long() string {
	return c.y
}

func (c Coord) Distance(_ Coord) float64 {
	return 0
}

func (c *Coord) String() string {
	return c.x + ", " + c.y
}
