package distance

import (
	"errors"
	"fmt"
	"lesson2/distance/point"
)

type PathEarth struct {
	points []point.Coordinate
}

func NewPathEarth(points []point.Coordinate) *PathEarth {
	return &PathEarth{points: points}
}

func (l *PathEarth) AddPoint(point point.Coordinate) *PathEarth {
	l.points = append(l.points, point)
	return l
}

func (l PathEarth) Distance() (distance float64, err error) {

	if len(l.points) < 2 {
		return distance, errors.New("кол-во точек меньше двух")
	}

	for i, p := range l.points {
		if i == len(l.points)-1 {
			break
		}
		distanceTwoPoint, _ := l.distanceTwoPoint(p, l.points[i+1])
		fmt.Println("Растояние между точками", p, l.points[i+1], " равно ", distanceTwoPoint)
		distance += distanceTwoPoint

	}
	return
}

func (l PathEarth) distanceTwoPoint(x point.Coordinate, y point.Coordinate) (result float64, err error) {
	line := NewLineEarth(x, y)
	result = line.Distance()
	return
}
