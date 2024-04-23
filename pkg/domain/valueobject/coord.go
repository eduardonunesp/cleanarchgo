package valueobject

type Coord struct {
	Lat  string
	Long string
}

func NewCoord(lat, long string) (Coord, error) {
	return Coord{
		Lat:  lat,
		Long: long,
	}, nil
}

func (c *Coord) String() string {
	return c.Lat + ", " + c.Long
}
