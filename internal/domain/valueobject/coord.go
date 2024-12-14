package valueobject

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)

var (
	ErrInvalidLatCoordinate  = errors.New("invalid latitude coordinate")
	ErrInvalidLongCoordinate = errors.New("invalid longitude coordinate")
	ErrFailedToParseLat      = errors.New("failed to parse latitude")
	ErrFailedToParseLong     = errors.New("failed to parse longitude")
)

const (
	minLat  = -90
	maxLat  = 90
	minLong = -180
	maxLong = 180
)

type Coordinate struct {
	lat  decimal.Decimal
	long decimal.Decimal
}

func BuildCoordinate(lat, long string) (Coordinate, error) {
	var (
		err error
		c   Coordinate
	)
	c.lat, err = decimal.NewFromString(lat)
	if err != nil {
		return Coordinate{}, ErrFailedToParseLat
	}
	c.long, err = decimal.NewFromString(long)
	if err != nil {
		return Coordinate{}, ErrFailedToParseLong
	}
	if c.lat.LessThan(decimal.NewFromInt(minLat)) || c.lat.GreaterThan(decimal.NewFromInt(maxLat)) {
		return Coordinate{}, ErrInvalidLatCoordinate
	}
	if c.long.LessThan(decimal.NewFromInt(minLong)) || c.long.GreaterThan(decimal.NewFromInt(maxLong)) {
		return Coordinate{}, ErrInvalidLongCoordinate
	}
	return c, nil
}

func (c Coordinate) Lat() string {
	return c.lat.String()
}

func (c Coordinate) Long() string {
	return c.long.String()
}

func (c Coordinate) Distance(d Coordinate) float64 {
	lat1, _ := strconv.ParseFloat(c.Lat(), 64)
	long1, _ := strconv.ParseFloat(c.Long(), 64)
	lat2, _ := strconv.ParseFloat(d.Lat(), 64)
	long2, _ := strconv.ParseFloat(d.Long(), 64)
	return haversine(lat1, long1, lat2, long2)
}

func haversine(lat1, long1, lat2, long2 float64) float64 {
	const earthRadius = 6371
	lat1Rad := toRadians(lat1)
	lat2Rad := toRadians(lat2)
	deltaLat := toRadians(lat2 - lat1)
	deltaLong := toRadians(long2 - long1)
	a := hav(deltaLat) + (hav(deltaLong) * hav(lat1Rad) * hav(lat2Rad))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c
}

func toRadians(deg float64) float64 {
	return deg * (3.14159265359 / 180)
}

func hav(theta float64) float64 {
	return (1 - math.Cos(theta)) / 2
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%s, %s", c.lat.String(), c.long.String())
}
