// Always have main package with main function defined inside
// Package = group of related code
package main

import "fmt"

const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "

// returns a string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Portuguese" {
		return portugueseHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("anything", "Xhosa"))
}
