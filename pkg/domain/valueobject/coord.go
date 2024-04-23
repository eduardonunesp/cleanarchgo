package valueobject

type Coord struct {
	lat  string
	long string
}

func NewCoord(lat, long string) (Coord, error) {
	return Coord{lat, long}, nil
}

func (c Coord) Lat() string {
	return c.lat
}

func (c Coord) Long() string {
	return c.long
}

func (c *Coord) String() string {
	return c.lat + ", " + c.long
}
