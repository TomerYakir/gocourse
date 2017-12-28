package main

import "fmt"

// this means that if you have a receiver called getGreeting() that returns a string, you are a bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// receiver value omitted as it's not being used
func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

/*
func (eb englishBot) printGreeting() {
	fmt.Println(eb.getGreeting())
}

func (sb spanishBot) printGreeting() {
	fmt.Println(sb.getGreeting())
}
*/
