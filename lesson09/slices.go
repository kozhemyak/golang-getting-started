package main

import "fmt"

// Slices
func basicSlice() {
	fmt.Println("---\nSlices:")

	// Basic array with 6 size
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Slice
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// slice[low : high]
	var s []int = primes[1:4]
	fmt.Println(s)
}

// Slices are like references to arrays
func sliceChangeArray() {
	fmt.Println("---\nSlices are like references to arrays:")
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.

	// Create and set the array
	names := [4]string{
		"Dmitriy",
		"Nikolay",
		"Evgeniy",
		"Maxim",
	}

	// Create and set slices
	// sliceA := names[0:2]
	sliceB := names[1:3]

	// Change slice and array at the same time
	sliceB[0] = "Yana"
	fmt.Printf("names: %v\n", names)
}

// Slice literals
func sliceLiteral() {
	fmt.Println("---\nSetting the slices:")

	// And this creates the same array as above, then builds a slice that references it:
	sliceLiteral := []bool{true, false, false, true}
	fmt.Printf("sliceLiteral: %v\n", sliceLiteral)

	structArray := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, true},
		{3, false},
		{4, true},
	}

	fmt.Printf("structArray: %v\n", structArray)
}

// Slice defaults
func sliceBounds() {
	fmt.Println("---\nSlice defaults and bounds:")

	// Ommiting the bounds
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

// Slice length and capacity
func sliceLenCap() {
	fmt.Println("---\nSlice length and capacity:")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("s", s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice("s", s)

	// Extend its length.
	s = s[:4]
	printSlice("s", s)

	// Drop its first two values.
	s = s[2:]
	printSlice("s", s)
}

// Nil slices
func sliceNil() {
	fmt.Println("---\nNil slices:")
	var s []int

	if s == nil {
		fmt.Println("The slice is nil!")
	} else {
		printSlice("s", s)
	}
}

// Creating a slice with make
func makingSlices() {
	fmt.Println("---\nCreating a slice with make:")

	a := make([]int, 5) // len(a)=5
	printSlice("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice("b", b)

	c := b[:cap(b)] // len(b)=5, cap(b)=5
	printSlice("c", c)

	d := c[2:5] // len(b)=3, cap(b)=3
	printSlice("d", d)
}

func main() {
	basicSlice()
	sliceChangeArray()
	sliceLiteral()
	sliceBounds()
	sliceLenCap()
	sliceNil()
	makingSlices()
}

func printSlice(name string, slice []int) {
	fmt.Printf("%s: length=%d capacity=%d of the slice:%v\n",
		name, len(slice), cap(slice), slice)
}
