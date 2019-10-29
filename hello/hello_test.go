// our hello func is part of main package
package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println("TEST NAME:", t.Name())

	// declaring variables
	got := Hello()
	want := "Hello Wild World!"

	if got != want {
		// Errorf "method" prints formatted message and fails test
		t.Errorf("got:%q, expected:%q", got, want)
	}
}
