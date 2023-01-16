package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b345",
		"white": "#ffff",
	}

	for color, hex := range colors {
		fmt.Println("Hex code for", color, "is", hex)
	}

}
