package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
	
		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("unknown")
		want := ""

		assertError(t, err, ErrNotFound)
		assertStrings(t, got, want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error)  {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}