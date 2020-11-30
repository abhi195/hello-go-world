package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	checklist := []string{
		"http://facebook.com",
		"http://amazon.com",
		"http://apple.com",
		"http://netflix.com",
		"http://google.com",
	}

	// creating channel for communicating between routines
	c := make(chan string)

	for _, link := range checklist {
		go checkStatus(link, c)
	}

	// iterating over channel in blocking manner
	for l := range c {
		// function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}

}

func checkStatus(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is unreachable!")
		c <- link
	}
	fmt.Println(link, "is up!")
	c <- link
}
