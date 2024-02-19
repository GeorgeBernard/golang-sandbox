package main

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

func rectangleInfo(r rectangle) {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

func makeSquare(r *rectangle) {
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}

func main() {
	r := rectangle{length: 4.2, width: 2.3}
	rectangleInfo(r)
	makeSquare(&r)
	rectangleInfo(r)
}
