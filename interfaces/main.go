package main

import (
	"fmt"
)

type bot interface {
	getGreetings() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreetings(sb)
	printGreetings(eb)
}

func (englishBot) getGreetings() string {
	return "Hi, There!"
}

func (spanishBot) getGreetings() string {
	return "Hola!"
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}
