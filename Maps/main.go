package main

import "fmt"

func main() {
	/*
		colors := map[string]string{ // map of keys of type string and values of type string
			"red":   "#ff0000",
			"green": "#ffdd00",
		}
	*/
	//var colors map[string]string
	colors := make(map[string]string) // same as above
	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"

	// delete(colors, "red") // delete a key

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
