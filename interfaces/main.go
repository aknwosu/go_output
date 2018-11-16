package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	sb := spanishBot{}
	eb := englishBot{}

	printGreeting(sb)
	printGreeting(eb)
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "hola"
}
