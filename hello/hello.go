// Always have main package with main function defined inside
// Package = group of related code
package main

import "fmt"

const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

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
	case "Spanish":
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("anything", "Xhosa"))
}
