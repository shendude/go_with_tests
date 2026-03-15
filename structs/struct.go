package structs_example

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) MutA() {
	r.Width *= 2
	r.Height *= 2
}

func (r *Rectangle) MutB() {
	r.Width *= 2
	r.Height *= 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	sideA float64
	sideB float64
}

func (t Triangle) Area() float64 {
	return Rectangle{t.sideA, t.sideB}.Area() / 2
}
