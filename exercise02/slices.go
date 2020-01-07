package main

import "golang.org/x/tour/pic"

// The task:
// https://tour.golang.org/moretypes/18
// Implement Pic. It should return a slice of length dy,
// each element of which is a slice of dx 8-bit unsigned integers.
// When you run the program, it will display your picture,
// interpreting the integers as grayscale (well, bluescale) values.

func Pic(dx, dy int) [][]uint8 {

	slice := make([][]uint8, dy)

	for i := range slice {
		subSlice := make([]uint8, dx)

		for index := 0; index < dx; index++ {
			subSlice[index] = uint8(rand.Intn(255))
		}

		slice[i] = subSlice
	}

	return slice
}

func main() {
	pic.Show(Pic)
}
