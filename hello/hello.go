// Always have main package with main function defined inside
// Package = group of related code
package main

import "fmt"

// returns a string
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("anything"))
}
