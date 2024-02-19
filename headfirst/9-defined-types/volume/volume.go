package main

import "fmt"

type Liters float64
type Millileters float64
type Gallons float64

func (l Liters) ToMilliliters() Millileters {
	return Millileters(l * 1000)
}

func (m Millileters) ToLiters() Liters {
	return Liters(m / 1000)
}

func main() {
	l := Liters(3)
	fmt.Printf("%0.1f liters is %0.1f milliliters\n", l, l.ToMilliliters())
	m := Millileters(500)
	fmt.Printf("%0.1f milliliters is %0.1f liters\n", m, m.ToLiters())
}
