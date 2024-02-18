package main

import (
	"fmt"
)

func main() {
	// Original
	// var pebbleweight float64 = 0.1
	// var rockweight float64 = 1.2
	// var boulderweight float64 = 502.4
	// var total_weight float64 = pebbleweight + rockweight + boulderweight
	// fmt.Println(total_weight)

	pebbleWeight, rockWeight, boulderWeight := 0.1, 1.2, 502.4
	totalWeight := pebbleWeight + rockWeight + boulderWeight
	fmt.Println(totalWeight)
}
