package main

import (
	"fmt"
	"math"
)

// Functions are values too.
// Function values may be used as function arguments and return values.

// (<FUNCNAME> func(input types) output type) output type { return <FUNCNAME> (VALUES) }
func compute(fn func(float64, float64) float64) float64 {
	// Pass the values to the declared fucntion above
	return fn(9, 13)
}

func main() {
	// They can be passed around just like other values.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// Invoke function with values
	fmt.Println(hypot(5, 12))

	// Invoke function that invokes another function passed through the argumetns
	fmt.Println(compute(hypot))

	// Invoke function that invokes math.Pow passed through the argumetns
	fmt.Println(compute(math.Pow))
}
