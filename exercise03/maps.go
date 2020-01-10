package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// Implement WordCount.
// https://tour.golang.org/moretypes/23
// It should return a map of the counts of each “word”
// in the string s. The wc.Test function runs a test
// suite against the provided function and prints
// success or failure.

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, v := range strings.Fields(s) {
		m[v] = m[v] + 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
