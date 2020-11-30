package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// greetings
	printGreeting(eb)
	printGreeting(sb)

	// bot versions
	printVersion(eb)
	printVersion(sb)

	// accessing google.com via http
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println("Google.com html : ")

	// using Stdout Writer
	// io.Copy(os.Stdout, resp.Body)

	// using custom Writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
