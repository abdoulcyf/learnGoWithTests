package arrays_and_slices

import (
	"fmt"
	"reflect"
	"testing"
)

var _ fmt.Formatter // for debugging only,delete when done

func TestSumTails(t *testing.T) {

	t.Run("return sum of all values except the tails", func(t *testing.T) {
		got := sumTails([]int{1, 2}, []int{5, 9})
		want := []int{2, 9}

		assertSumATails(t, got, want)
	})

	//-------------------------------------------------------
	t.Run("Safely sum empy slices", func(t *testing.T) {
		got := sumTails([]int{}, []int{5, 9})
		want := []int{0, 9}

		assertSumATails(t, got, want)
	})
}

func assertSumATails(t testing.TB, got, want []int) {
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
