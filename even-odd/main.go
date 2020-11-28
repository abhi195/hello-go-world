package main

import "fmt"

func main() {
	slice := []int{}

	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}

	for _, val := range slice {
		if val%2 == 0 {
			fmt.Println(val, "is even")
		} else {
			fmt.Println(val, "is odd")
		}
	}
}
