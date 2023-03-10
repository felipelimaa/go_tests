package geometrics

import (
	"math"
)

// Form
type Form interface {
	Area() float64
}

// Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle
type Circle struct {
	Radius float64
}

// Area
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
