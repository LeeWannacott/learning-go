package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	// var colors map[int]string
	// colors := make(map[int]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ff0000",
		"white": "#ffffff",
	}
	// colors[39] = "#ffffff"
	// delete(colors, 39)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}

}
