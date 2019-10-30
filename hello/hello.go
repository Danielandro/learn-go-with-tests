// Always have main package with main function defined inside
// Package = group of related code
package main

import "fmt"

const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, "

// returns a string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "Portuguese":
		prefix = portugueseHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("anything", "Xhosa"))
}
