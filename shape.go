package golang_united_school_homework

import "math"

type Shape interface {
	// CalcPerimeter returns calculation result of perimeter
	CalcPerimeter() float64
	// CalcArea returns calculation result of area
	CalcArea() float64
}

func (c Circle) CalcPerimeter() float64 {
	return math.Pi * c.Radius * 2
}

func (c Circle) CalcArea() float64 {
	return math.Pi * c.Radius * c.Radius
}
