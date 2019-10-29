// our hello func is part of main package
package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	// this is a Subtest
	t.Run("saying hello to people", func(t *testing.T) {
		fmt.Println("TEST NAME:", t.Name())

		// declaring variables
		got := Hello("Ade")
		want := "Hello, Ade"

		if got != want {
			// Errorf "method" prints formatted message and fails test
			t.Errorf("got:%q, expected:%q", got, want)
		}
	})

	t.Run("say 'Hello, World' when empty string supplied", func(t *testing.T) {
		got := Hello()
		want := "Hello, World"

		if got != want {
			t.Errorf("Expected: %q Got: %q", want, got)
		}
	})
}
