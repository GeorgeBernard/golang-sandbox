package main

import "fmt"

func main() {
	days := [7]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	for i, day := range days {
		fmt.Println(i, day)
	}
}
