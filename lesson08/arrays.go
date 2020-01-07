package main

import "fmt"

func main() {

	// Fixed-sise arrays, cannot be resized
	var array [5]string

	array[0] = "Happy"
	array[1] = "New"
	array[2] = "Year"
	array[3] = "Dear"
	array[4] = "Fellas"

	fmt.Println(array)

	// Another one
	primes := [10]int{2, 3, 4, 5, 6}
	fmt.Println(primes)
	fmt.Println(len(primes))
}
