package distance

import (
	"lesson2/distance/point"
	"math"
)

type Line2d struct {
	point1 point.Point2d
	point2 point.Point2d
}

func (l Line2d) Point1() point.Point2d {
	return l.point1
}

func (l Line2d) Point2() point.Point2d {
	return l.point2
}

func NewLine(point1 point.Point2d, point2 point.Point2d) *Line2d {
	return &Line2d{point1: point1, point2: point2}
}

func (l Line2d) Distance() (distance float64) {
	distance = math.Sqrt(math.Pow(l.point2.X()-l.point1.X(), 2) + math.Pow(l.point2.Y()-l.point1.Y(), 2))
	return
}
