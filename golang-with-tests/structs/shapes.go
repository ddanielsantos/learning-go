package structs

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Shape interface { // interface resolution is implicit
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
