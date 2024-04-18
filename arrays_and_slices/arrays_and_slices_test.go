package arrays_and_slices

import (
	"fmt"
	"testing"
)

var _ fmt.Formatter // for debugging only,delete when done

func TestSum(t *testing.T) {
	t.Skip()
	//---------------------------------------------
	t.Run("return collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 5, 9}
		got := Sum(nums)
		want := 20

		assertCorrectMessage(t, got, want, nums)
	})

	//---------------------------------------------
	t.Run("return collections of 6", func(t *testing.T) {
		nums := []int{1, 2, 3, 5, 9, 10}
		got := Sum(nums)
		want := 30

		assertCorrectMessage(t, got, want, nums)
	})

}

func assertCorrectMessage(t testing.TB, got, want int, nums []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but want %d given, %v", got, want, nums)
	}
}

// func ExampleSum() {
//   nums := []int{4, 5, 9}
//   char := Sum(nums)
//   fmt.Println(char)
//   // Output: 18
// }

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{4, 5, 9, 10}
		Sum(nums)
	}
}
