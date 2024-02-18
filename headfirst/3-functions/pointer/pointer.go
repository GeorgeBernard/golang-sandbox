package main

import "fmt"

func main() {
	var myInt int = 2
	var myIntPointer *int = &myInt
	fmt.Println(myIntPointer)
}
