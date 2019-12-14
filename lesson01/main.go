package main

import (
	"fmt"
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

// Named return values
func namedReturn(sum int) (x, y int) {
	// function level statement "var"
	var maximumOne int = 15

	// the := short assignment statement, same
	maximumTwo := 15

	sum = maximumOne
	sum = maximumTwo

	x, y = (5 + sum), (10 + sum)

	// "naked" return
	return
}

// Run all this calculations
func Start() {
	// First call
	fmt.Println("Hello World!", "Are you ready to \"Go\"?")

	// Function 1: add
	fmt.Println(add(13, 13))

	// Function 2: multipleVars
	a, b := multipleVars("One", "Two")
	fmt.Println(a, b)

	// Function 3:
	fmt.Println(namedReturn(1986))
}

func main() {
	Start()
}
