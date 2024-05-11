package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	//-------------------------------------------------------
	t.Run("Known Word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertSearch(t, got, want)
	})

	//-------------------------------------------------------
	t.Run("Unknown Word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("Unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertSearch(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	key := "test"
	value := "this is just a test"
	err := dictionary.Add(key, value)
	if err != nil {
		return
	}

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertSearch(t, got, want)
}

func assertSearch(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertNotFound(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Did not get expected error. Wanted %s", err)
	}

}
