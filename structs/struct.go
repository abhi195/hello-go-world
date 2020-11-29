package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

// pointer receiver
func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}

// value receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
