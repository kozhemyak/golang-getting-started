package main

import (
	"fmt"
	"strings"
)

// Slices of slices / Appending to a slice
func ticTacToe() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{" - ", " - ", " - "},
		[]string{" - ", " - ", " - "},
	}

	// The slice grows as needed.
	board = append(board,
		// We can add more than one element at a time.
		append([]string{" - "}, " - ", " - "))

	// The players take moves.
	board[0][0] = " X "
	board[2][2] = " O "
	board[1][2] = " X "
	board[1][0] = " O "
	board[0][2] = " X "

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// Ranges
func rangeTest(slice []int) {
	fmt.Println("Iterating the slice through the range:")

	// You can skip the index or value by assigning to _.
	// for i, _ := range pow
	// for _, value := range pow	- ForEach() like

	for i, v := range slice {
		fmt.Printf("\t%d: %d\n", i, v)

		// does not affect the original value
		v = 1e10
	}
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	rangeTest(pow)
}
