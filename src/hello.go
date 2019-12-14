package main

import (
	"fmt"
	//"math"
)

// package level statement "var"
var maxCount, minCount int
var isVisible, isCreated bool = true, false

// take the type of the initializer
var tryCount, isDone, preciseNumber = 13, false, 3.14

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

func main() {
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
