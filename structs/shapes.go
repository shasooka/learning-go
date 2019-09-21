package structs

import "math"

// Rectangle Struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter of the square
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area of the square
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle Struct
type Circle struct {
	Radius float64
}

// Perimeter of the circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle Struct
type Triangle struct {
	Height float64
	Base   float64
}

// Area of the triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Height * t.Base
}

// Perimeter of the circle
func (t Triangle) Perimeter() float64 {
	return 3 * t.Base
}

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}
