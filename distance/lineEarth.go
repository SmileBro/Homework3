package distance

import (
	"lesson2/distance/point"
	"math"
)

type LineEarth struct {
	coordinate1 point.Coordinate
	coordinate2 point.Coordinate
}

func (l LineEarth) Coordinate1() point.Coordinate {
	return l.coordinate1
}

func (l LineEarth) Coordinate2() point.Coordinate {
	return l.coordinate2
}

func NewLineEarth(coordinate1 point.Coordinate, coordinate2 point.Coordinate) *LineEarth {
	return &LineEarth{coordinate1: coordinate1, coordinate2: coordinate2}
}

func (l LineEarth) Distance() (result float64) {
	//Find the value of the latitude in radians:
	lat1 := l.coordinate1.Latitude() / 57.29577951
	lat2 := l.coordinate2.Latitude() / 57.29577951
	//Find the value of longitude in radians:
	long1 := l.coordinate1.Longitude() / 57.29577951
	long2 := l.coordinate2.Longitude() / 57.29577951
	//Distance, d = 3963.0 * arccos[(sin(lat1) * sin(lat2)) + cos(lat1) * cos(lat2) * cos(long2 – long1)]
	result = 3963.0 * math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(long2-long1)) * 1.609344
	return
}
