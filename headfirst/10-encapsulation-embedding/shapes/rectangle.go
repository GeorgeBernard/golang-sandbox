package shapes

import (
	"fmt"
)

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Length() float64 {
	return r.length
}

func (r *Rectangle) SetLength(length float64) error {
	if length < 0 {
		return fmt.Errorf("invalid length: %f", length)
	}
	r.length = length
	return nil
}

func (r *Rectangle) Width() float64 {
	return r.width
}

func (r *Rectangle) SetWidth(width float64) error {
	if width < 0 {
		return fmt.Errorf("invalid width: %f", width)
	}
	r.width = width
	return nil
}

func (r *Rectangle) Info() {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

// Convert this function to a method on the "rectangle" type.
func (r *Rectangle) MakeSquare() {
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}
