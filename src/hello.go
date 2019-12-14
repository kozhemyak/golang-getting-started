package main

import (
	"fmt"
	//"math"
)

// my first function
func add(x int, y int) int {
	return x + y
}

// omit the type for two input variables
func multipleVars(x, y string) (string, string) {
	// return multiple results
	return x, y
}

func main() {
	// First call
	fmt.Println("Hello World!", "Are you ready to \"Go\"?")

	// Function 1 call
	fmt.Println(add(13, 13))

	// Function 2 call
	a, b := multipleVars("One", "Two")
	fmt.Println(a, b)
}
