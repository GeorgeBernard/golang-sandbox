package main

import "fmt"

type pet struct {
	name string
	age  int
}

func main() {
	var max pet

	max.name = "Max"
	max.age = 5
	fmt.Println("Name:", max.name)
	fmt.Println("Age:", max.age)
}
