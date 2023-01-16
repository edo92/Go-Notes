package main

import "fmt"

func main() {
	// Slice value passed by reference
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
