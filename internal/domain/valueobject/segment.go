package valueobject

import "fmt"

type Segment struct {
	from Coordinate
	to   Coordinate
}

func BuildSegment(from Coordinate, to Coordinate) Segment {
	return Segment{from: from, to: to}
}

func BuildSegmentFromCoordinates(fromLat, fromLong, toLat, toLong string) (Segment, error) {
	from, err := BuildCoordinate(fromLat, fromLong)
	if err != nil {
		return Segment{}, err
	}
	to, err := BuildCoordinate(toLat, toLong)
	if err != nil {
		return Segment{}, err
	}
	return BuildSegment(from, to), nil
}

func (s Segment) From() Coordinate {
	return s.from
}

func (s Segment) To() Coordinate {
	return s.to
}

func (s Segment) Length() float64 {
	return s.from.Distance(s.to)
}

func (s Segment) String() string {
	return fmt.Sprintf("%s -> %s", s.from, s.to)
}
