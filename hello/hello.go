// Always have main package with main function defined inside
// Package = group of related code
package main

import "fmt"

const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "

// PUBLIC FUNCTION (first-letter uppercase)
// returns a string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// PRIVATE FUNCTION (first-letter lowercase)
// prefix is a named return value (assigned the value "" aka zero value)
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Portuguese":
		prefix = portugueseHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// returns prefix implicitely
	return
}

func main() {
	fmt.Println(Hello("anything", "Xhosa"))
}
