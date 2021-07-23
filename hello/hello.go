package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

func Hello(s, l string) string {
	if s == "" {
		s = "World"
	}
	return greetingPrefix(l) + s
}

func greetingPrefix(l string) (prefix string) {
	switch l {
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
	fmt.Println(Hello("", ""))
}
