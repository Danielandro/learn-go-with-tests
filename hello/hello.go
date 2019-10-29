// Always have main package with main function defined inside
// Package = group of related code
package main

import "fmt"

const englishHelloPrefix = "Hello, "

// returns a string
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("anything"))
}
