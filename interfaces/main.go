package main

import "fmt"

type englisBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting(int) string
}

func (eb englisBot) getGreeting(i int) string {
	return "Hello"
}
func (eb spanishBot) getGreeting(i int) string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting(1))
}

func main() {
	eb := englisBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
