package main

import (
	"fmt"
)

func main() {
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// }

	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["white"] = "#FFF"
	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff000",
		"green": "#EDEE",
		"white": "#FFF",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
