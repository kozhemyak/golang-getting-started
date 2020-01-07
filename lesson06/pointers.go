package main

import (
	"fmt"
	"time"
)

// Learn Pointers
// default zero value is nil
// *T is pointer to a T value in memory
// & operator generates a pointer to its operand
// * operator denotes the pointer's value
// deferencing / indirecting
// Unlike C, Go has no pointer arithmetic.

func main() {
	answerToLife, bornYear := 21, 1986

	// point to answerToLife
	pointer := &answerToLife

	// set answerToLife through the pointer
	*pointer = 42

	// point to bornYear
	pointer = &bornYear

	// calculate and set the bornYear through the pointer
	*pointer = time.Now().Year() - int(*pointer)

	fmt.Printf("The answer to life %v is stored in %v memmory address\n", answerToLife, &answerToLife)
	fmt.Printf("You are %v years old and the information about it is stored here: %v\n", *pointer, pointer)
}
