// our hello func is part of main package
package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // marks this as a helper function
		if got != want {
			t.Errorf("Expected: %q Got: %q", want, got)
		}

	}
	// this is a Subtest
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ade", "")
		want := "Hello, Ade"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when empty string supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Andreas", "Portuguese")
		want := "Olá, Andreas"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Francois", "French")
		want := "Bonjour, Francois"

		assertCorrectMessage(t, got, want)
	})
}
