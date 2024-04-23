package valueobject

type Segment struct {
	from Coord
	to   Coord
}

func BuildSegment(from Coord, to Coord) Segment {
	return Segment{from: from, to: to}
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
