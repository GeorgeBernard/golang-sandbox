package main

import "fmt"

func main() {
	// We'll count the number of times each number occurs
	// within this array.

	// numbers := [9]int{1, 0, 2, 0, 1, 0, 0, 1, 2}
	numbers := [100]int{4, 2, 6, 7, 8, 0, 1, 8, 7, 8,
		1, 5, 3, 2, 2, 1, 9, 6, 1, 0, 0, 0, 8, 4, 6,
		2, 2, 4, 7, 9, 6, 5, 9, 0, 5, 1, 1, 5, 4, 7,
		7, 9, 7, 8, 6, 3, 3, 3, 4, 8, 0, 4, 1, 1, 7,
		9, 8, 8, 1, 2, 3, 6, 4, 9, 2, 5, 8, 6, 7, 7,
		5, 4, 2, 9, 4, 4, 2, 2, 5, 5, 0, 0, 0, 9, 1,
		9, 5, 8, 0, 1, 1, 0, 5, 3, 8, 6, 3, 4, 4, 9}

	// YOUR CODE HERE: Process each element of "numbers",
	// keeping track of how many times you've seen 0, 1,
	// or 2.
	// Finally, print out how many times each number
	// occurred.

	var max int = 0
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	counts := make([]int, max+1)
	for _, n := range numbers {
		counts[n]++
	}

	for i, count := range counts {
		fmt.Println(i, "occurred", count, "times")
	}
}
