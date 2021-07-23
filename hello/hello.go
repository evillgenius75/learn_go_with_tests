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

	if l == spanish {
		return spanishHelloPrefix + s
	} else if l == french {
		return frenchHelloPrefix + s
	}
	return englishHelloPrefix + s
}

func main() {
	fmt.Println(Hello("", ""))
}
