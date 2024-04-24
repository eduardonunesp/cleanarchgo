package valueobject

type Segment struct {
	from Coord
	to   Coord
}

func BuildSegment(from Coord, to Coord) Segment {
	return Segment{from: from, to: to}
}

func BuildSegmentFromCoords(fromLat, fromLong, toLat, toLong string) (Segment, error) {
	from, err := BuildCoord(fromLat, fromLong)
	if err != nil {
		return Segment{}, err
	}
	to, err := BuildCoord(toLat, toLong)
	if err != nil {
		return Segment{}, err
	}
	return BuildSegment(from, to), nil
}

func (s Segment) From() Coord {
	return s.from
}

func (s Segment) To() Coord {
	return s.to
}

func (s Segment) Length() float64 {
	return s.from.Distance(s.to)
}

func (s Segment) String() string {
	return s.from.String() + " -> " + s.to.String()
}
