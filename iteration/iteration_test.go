package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	//---------------------------------------------
	t.Run("repeat char a 6 times", func(t *testing.T) {
		got := Repeat('a', 6)
		want := "aaaaaa"

		assertCorrectMessage(t, got, want)
	})

	//---------------------------------------------------
	t.Run("if char == 'A', repeat 'a' 3 times", func(t *testing.T) {
		got := Repeat('A', 3)
		want := "aaa"

		assertCorrectMessage(t, got, want)
	})

	//---------------------------------------------------
	t.Run("if char == 'B', repeat 'b'", func(t *testing.T) {
		got := Repeat('B', 10)
		want := "bbbbbbbbbb"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func ExampleRepeat() {
	char := Repeat('a', 2)
	fmt.Println(char)
	// Output: aa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat('c', 2)
    Repeat('D', 8)
	}
}
