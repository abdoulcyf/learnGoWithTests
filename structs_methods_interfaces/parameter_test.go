package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("return parama of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 5.0}
		got := rectangle.Perimeter()
		want := 30.0

		assertPerimeter(t, got, want)
	})
}

// func TestArea(t *testing.T){
// 	t.Run("return area of rectangle", func(t *testing.T){
// 		rectangle := Rectangle{10.0, 10.0}
// 		got := rectangle.Area() // Fix: Pass two arguments to the Area function
// 		want := 100.0

// 		if got != want {
// 			t.Errorf("got: %.2f, want: %.2f", got, want)
// 		}
// 	})

// 	//---------------------------------------------------------
// 	t.Run("#3return area of Circle", func(t *testing.T){
// 		circle := Circle{10.0}
// 		got := circle.Area() // Fix: Pass two arguments to the Area function
// 		want := 314.1592653589793

// 		assertArea(t, got, want)
// 	})
// }

func assertPerimeter(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func assertArea(t testing.TB, got, want float64) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("got %g want %g", got, want)
// 	}
// }

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("return area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100)
	})

	//---------------------------------------------------------
	t.Run("#3return area of Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
