package main

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@email.com",
			zipCode: 12345,
		},
	}

	// go shortcut, auto "&" pointer referencing
	alex.updateFirstName("Alexa")
	alex.print()
}
