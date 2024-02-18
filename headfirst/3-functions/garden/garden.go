package main

import (
	"fmt"
	"log"
)

// YOUR CODE HERE: Write a "perimeter" function.
func perimeter(length float64, width float64) (float64, error) {
	if length < 0 {
		return 0, fmt.Errorf("a length of %.2f is invalid", length)
	} else if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	} else {
		return 2*length + 2*width, nil
	}
}

func main() {
	// YOUR CODE HERE: Call "perimeter" three times.
	// Add the three return values together, and store the
	// total in the "total" variable.

	total, err := perimeter(-1, -2)
	if err != nil {
		log.Fatal(err)
	}

	// total := perimeter(8.2, 10)
	// total += perimeter(5, 5.4)
	// total += perimeter(6.2, 4.5)

	fmt.Println("You'll need", total, "meters of fencing")
}
