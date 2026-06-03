package main

import (
	"fmt"
	"strings"
)


const (
	spanish = "spanish"
	french  = "french"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	language = strings.ToLower(language)
	prefix := greetingPrefix(language)
	return fmt.Sprintf("%s%s", prefix, name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}