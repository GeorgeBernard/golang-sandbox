package main

import "fmt"

// YOUR CODE HERE: Define a printLines function.

func main() {
	daysOfWeek := [7]string{"Sunday", "Monday", "Tuesday",
		"Wednesday", "Thursday", "Friday", "Saturday"}

	// YOUR CODE HERE: Get a slice containing just the
	// weekdays.
	weekDays := daysOfWeek[1:6]
	// Pass that slice to printLines.
	for _, day := range weekDays {
		fmt.Println(day)
	}
}
