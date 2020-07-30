package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to Toks", func(t *testing.T) {
		got := Hello("Ige", "Adetokunbo", "french")
		want := "Bonjour, Ige Adetokunbo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World People'", func(t *testing.T) {
		got := Hello("", "", "")
		want := "Hello, World People"
		assertCorrectMessage(t, got, want)
	})
}
