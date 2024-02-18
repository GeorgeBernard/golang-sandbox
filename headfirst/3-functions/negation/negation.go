package main

import "fmt"

func negate(b *bool) {
	*b = !*b
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}
