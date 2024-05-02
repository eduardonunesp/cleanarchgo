package valueobject

type Coord struct {
	x string
	y string
}

func BuildCoord(lat, long string) (Coord, error) {
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
