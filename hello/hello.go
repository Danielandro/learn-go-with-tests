// Always have main package with main function defined inside
// Package = group of related code
package main

import "fmt"

// returns a string
func Hello() string {
	return "Hello Wild World!"
}

func main() {
	fmt.Println(Hello())
}
