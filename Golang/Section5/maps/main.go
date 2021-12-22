package main

import (
	"fmt"
)

func main() {
	// METHOD1
	// map with all the keys of type string, and all the values of type string
	// colors := map[string]string{
	// 	"red": "#FF0000",
	// 	"green": "#00FF00",
	// }

	// METHOD2
	// var colors map[string]string

	// METHOD3
	// colors := make(map[string]string)
	// colors["yellow"] = "#FF01111"

	// delete(colors, "yellow")

	colors := map[string]string{
		"red": "#FF0000",
		"green": "#00FF00",
		"white": "#00FF01",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, color := range c {
		fmt.Println(key, color)
	}
}