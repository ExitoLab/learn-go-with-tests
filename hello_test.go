package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello to Toks", func(t *testing.T) {
		got := Hello("Ige", "Adetokunbo")
		want := "Hello, Ige Adetokunbo"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World People"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

}
