package arrays_and_slices

import (
	"fmt"
	"reflect"
	"testing"
)

var _ fmt.Formatter // for debugging only,delete when done

func TestSumAll(t *testing.T) {
	t.Run("return sum of all values", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertSumAll(t, got, want)
	})
	//-------------------------------------------------------
	
	t.Run("return sum of all values except the tails", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{5, 9})
		want := []int{2, 9}

		assertSumAll(t, got, want)
	})

}

func assertSumAll(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// func ExampleSumAll() {
// 	sum := SumAll([]int{7, 9}, []int{9, 7})
// 	fmt.Println(sum)
// 	// Output: []int{16, 16}
// }
