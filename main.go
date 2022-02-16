package main

import (
	"fmt"
	"lesson2/distance"
	"lesson2/distance/navigator"
	"lesson2/distance/point"
)

func main() {
	//New York City Latitude and Longitude
	nyLat := 40.730610
	nyLong := -73.935242
	//Los Angeles Latitude and Longitude
	laLat := 34.052235
	laLong := -118.243683
	line := distance.NewLineEarth(*point.NewCoordinate(nyLat, nyLong), *point.NewCoordinate(laLat, laLong))
	//The total straight line distance between New York and Los Angeles is 3936 KM
	fmt.Println(line.Distance())
	pathEarth := distance.NewPathEarth(make([]point.Coordinate, 0))
	//The total straight line distance between Los Angeles and Washington Dc is 3693 KM
	waLat := 38.889248
	waLong := -77.050636
	pathEarth.AddPoint(*point.NewCoordinate(nyLat, nyLong)).AddPoint(*point.NewCoordinate(laLat, laLong))
	pathEarth.AddPoint(*point.NewCoordinate(waLat, waLong))

	navi := navigator.NewNavigator([]navigator.AllDistance{pathEarth})

	fullPath, err := navi.Path()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Полный путь по навигатору %v\n", fullPath)

}
