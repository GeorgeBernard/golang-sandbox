package main

import (
	"fmt"
	"log"
	"sandbox/headfirst/10-encapsulation-embedding/geo"
	"sandbox/headfirst/10-encapsulation-embedding/shapes"
)

func testCoordinates() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testLandmark() {
	location := geo.Landmark{}
	err := location.SetName("The Googleplex")
	check(err)
	err = location.SetLatitude(37.42)
	check(err)
	err = location.SetLongitude(-122.08)
	check(err)
	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}

func testShapes() {
	var r shapes.Rectangle
	err := r.SetLength(4.2)
	check(err)
	fmt.Println("Length:", r.Length())
	err = r.SetWidth(-2.3)
	check(err)
	fmt.Println("Width:", r.Width())
}

func main() {
	//testCoordinates()
	//testLandmark()
	testShapes()
}
