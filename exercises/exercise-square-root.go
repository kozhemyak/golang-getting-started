package main

import (
	"fmt"
	"math"
)

// The task:
// let's implement a square root function:
// 	given a number x, we want to find the number z
//	for which z² is most nearly x.
//	"y^2 = x" need to find "y"

func sqrt(x float64) float64 {

	if x < 0 {
		return math.NaN()
	}

	if (x == 0) || (x == 1) {
		return x
	}

	const minIteration = 5
	var squareRoot float64 = x / 2

	for i := 0; i <= minIteration; i++ {
		squareRoot = math.Round(math.Abs(((squareRoot - (squareRoot * squareRoot) - x) / (2 * squareRoot))))
		//fmt.Printf("\t The new possible square root is %v\n", squareRoot)
	}

	return squareRoot
}

func main() {
	var number float64

	number = -1
	fmt.Printf("Start seeking the square root for %v\n", number)
	fmt.Println(sqrt(number))
	fmt.Println(math.Sqrt(number))

	number = 0
	fmt.Printf("Start seeking the square root for %v\n", number)
	fmt.Println(sqrt(number))
	fmt.Println(math.Sqrt(number))

	number = 1
	fmt.Printf("Start seeking the square root for %v\n", number)
	fmt.Println(sqrt(number))
	fmt.Println(math.Sqrt(number))

	number = 2
	fmt.Printf("Start seeking the square root for %v\n", number)
	fmt.Println(sqrt(number))
	fmt.Println(math.Sqrt(number))

	number = 4
	fmt.Printf("Start seeking the square root for %v\n", number)
	fmt.Println(sqrt(number))
	fmt.Println(math.Sqrt(number))

	number = 64
	fmt.Printf("Start seeking the square root for %v\n", number)
	fmt.Println(sqrt(number))
	fmt.Println(math.Sqrt(number))
}
