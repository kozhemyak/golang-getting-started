package main

import (
	"fmt"
)

// Read: https://blog.golang.org/defer-panic-and-recover

func defferedFunction01() {
	// The deferred call's arguments are evaluated immediately,
	// but the function call is not executed until the surrounding
	// function returns.
	defer fmt.Println("New Year!")
	fmt.Println("Happy")
}

// Stacking defers
func defferedFunction02() {
	fmt.Println("counting start...")

	for i := 0; i < 10; i++ {
		// Deferred function calls are pushed onto a stack.
		// its deferred calls are executed in last-in-first-out order.
		// deffered 0 -> 10, executed 10 -> 0
		defer fmt.Printf("\tdeffered %d\n", i)
	}

	fmt.Println("counting done.")
}

func main() {
	defferedFunction01()
	defferedFunction02()
}
