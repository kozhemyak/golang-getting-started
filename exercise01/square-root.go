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
	const minIteration = 10
	var squareRoot float64 = x / 2

	if x < 0 {
		return math.NaN()
	}

	if (x == 0) || (x == 1) {
		return x
	}

	for i := 0; i <= minIteration; i++ {
		squareRoot = math.Round(math.Abs(((squareRoot - (squareRoot * squareRoot) - x) / (2 * squareRoot))))
	}

	return squareRoot
}

func main() {
	numberSlice := []float64{-1, 0, 1, 2, 4, 8, 16, 32, 4096}

	for _, element := range numberSlice {
		fmt.Printf("Start seeking the square root for %v\n", element)
		fmt.Println(sqrt(element))
		fmt.Println(math.Sqrt(element))
	}
}
