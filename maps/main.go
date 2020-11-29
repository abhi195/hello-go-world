package main

import "fmt"

func main() {

	// init
	// 1
	// colors := make(map[string]string)
	// 2
	colors := map[string]string{
		"red":    "#FF0000",
		"yellow": "#FFFF00",
		"white":  "#FFFFFF",
	}

	// key add/update
	colors["black"] = "#000000"

	// key deletion
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(m map[string]string) {
	// iterating over map
	for key, value := range m {
		fmt.Println("Hex code for color", key, "is", value)
	}
}
