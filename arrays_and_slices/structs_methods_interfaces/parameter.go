package structs_methods_interfaces

type Rectangle struct {
	Width float64
	Height float64
}

const ParamMultiplier = 2

func (r  Rectangle) Perimeter() float64 {
		perim := ParamMultiplier * (r.Height + r.Width)
		return perim
}	

func (r Rectangle)Area() float64 {
	area := r.Width * r.Height
	return area
}
