package structs_methods_interfaces

import (
	//"fmt"
	"math"
)

//import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	R float64
}

const ParamMultiplier = 2

//const PI = 3.141592653589793238462643383279502884197

func (r Rectangle) Perimeter() float64 {
	perim := ParamMultiplier * (r.Height + r.Width)
	return perim
}

func (r Rectangle) Area() float64 {
	area := r.Width * r.Height
	return area
}

func (c Circle) Area() float64 {
	area := math.Pi * math.Pow(c.R, 2)
	return area
}

func PrintShape(r, w, h float64) float64 {
	var result float64
	shapes := []Shape{Circle{r}, Rectangle{Width: w, Height: h}}
	for _, shape := range shapes {
		result = shape.Area()
	}
	return result
}
