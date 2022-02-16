package point

type Point2d struct {
	x float64
	y float64
}

func (p Point2d) X() float64 {
	return p.x
}

func (p Point2d) Y() float64 {
	return p.y
}

func NewPoint2d(x float64, y float64) *Point2d {
	return &Point2d{x: x, y: y}
}
