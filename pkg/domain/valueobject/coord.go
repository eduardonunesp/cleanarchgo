package valueobject

type Coord struct {
	lat  string
	long string
}

func BuildCoord(lat, long string) (Coord, error) {
	return Coord{lat, long}, nil
}

func (c Coord) Lat() string {
	return c.lat
}

func (c Coord) Long() string {
	return c.long
}

func (c Coord) Distance(other Coord) float64 {
	return 0
}

func (c *Coord) String() string {
	return c.lat + ", " + c.long
}
