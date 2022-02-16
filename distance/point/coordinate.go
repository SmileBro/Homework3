package point

type Coordinate struct {
	latitude  float64
	longitude float64
}

func (c Coordinate) Latitude() float64 {
	return c.latitude
}

func (c Coordinate) Longitude() float64 {
	return c.longitude
}

func NewCoordinate(latitude float64, longitude float64) *Coordinate {
	return &Coordinate{latitude: latitude, longitude: longitude}
}
