package main

import "fmt"

type bot interface {
	getGreeting() string
	getVersion() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi, there!"
}

func (englishBot) getVersion() string {
	return "1.1.0"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (spanishBot) getVersion() string {
	return "1.0.0"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printVersion(b bot) {
	fmt.Println(b.getVersion())
}
