package structs_methods_interfaces

import "testing"


func TestPerimeter(t *testing.T){
	t.Run("return parama of a rectangle", func(t *testing.T){	
		rectangle := Rectangle{10.0, 5.0}
		got :=  rectangle.Perimeter()
		want := 30.0

		assertPerimeter(t, got, want)
	})
}

func TestArea(t *testing.T){
	t.Run("return area of rectangle", func(t *testing.T){
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area() // Fix: Pass two arguments to the Area function
		want := 100.0

		if got != want {
			t.Errorf("got: %.2f, want: %.2f", got, want)
		}
	})
}

func Area(height, width float64) float64 { // Fix: Modify the function signature to accept two float64 arguments
	panic("unimplemented")
}

func assertPerimeter(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

