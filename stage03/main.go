package main

import (
	"fmt"
)

func main() {
	sum := 0

	// For loop
	fmt.Println("For loop (init, cond, post)")

	for i := 0; i < 10; i++ {

		fmt.Printf("\t Add %v to sum (%v)\n", i, sum)
		sum += i
	}

	// For loop continued
	// The init and post statements are optional.
	fmt.Println("For loop (only cond):")

	for sum < 1000 {
		iteration := 300
		fmt.Printf("\t Add %v to sum (%v)\n", iteration, sum)
		sum += iteration
	}

	// Forever loop
	// The same as while in other languages
	fmt.Println("For loop (infinity):")
	for {
		iteration := 500
		fmt.Printf("\t Minus %v to sum (%v)\n", iteration, sum)
		sum -= 500

		if sum < 0 {
			break
		}
	}

	fmt.Println("--------------------")
	fmt.Printf("The total sum is %v \n\n", sum)

	////////////////////////////////////////////////////////////////
	// IF STATMENTS
	// simple
	if sum == -255 {
		fmt.Println("After some \"very\" coplex logic th sum equel -255")
	} else {
		fmt.Println("What heck is going on here?! Are they nuts?")
	}

	if step := 1000; sum > 0 {
		nextSum := sum + step

		fmt.Printf("The sum %v with step %v equel %v \n\n", sum, step, nextSum)
	}

}
