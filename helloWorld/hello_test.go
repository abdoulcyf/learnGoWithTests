package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say  Hello", func(t *testing.T) {
		got := Hello("Abdoul", "English")
		want := "Hello Abdoul"

		assertCorrectMessage(t, got, want)
	})

	//------------------------------------
	t.Run("Say hello Abdoul when arg is empty ", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello Abdoul"

		assertCorrectMessage(t, got, want)
	})

	//------------------------------------
	t.Run("Say hello Hola Abdoul if spanish ", func(t *testing.T) {
		got := Hello("Abdoul", "Spanish")
		want := "Hola Abdoul"

		assertCorrectMessage(t, got, want)
	})

	//------------------------------------
	t.Run("Say hello Bonjour Abdoul if French ", func(t *testing.T) {
		got := Hello("Abdoul", "French")
		want := "Bonjour Abdoul"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
