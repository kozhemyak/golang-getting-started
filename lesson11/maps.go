package main

import "fmt"

// struct type
type Vertex struct {
	Lat, Long float64
}

// ----------------------------------------------
// Maps
func maps() {
	// map [<NAME_TYPE>]<VALUE_TYPE>
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Nikolay"] = Vertex{40.1, -74.3}

	fmt.Println(m)
}

// ----------------------------------------------
// Map literals
func mapLiterals() {
	var m = map[string]Vertex{
		"Nikolay": Vertex{40.1, -74.3},
		"Dmitriy": Vertex{50.1, -10.3},
	}

	fmt.Println(m)
}

// ----------------------------------------------
// Map literals continued
func mapLiteralsContinued() {
	// ommiting the type if the top-level type is the same
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

// ----------------------------------------------
// Map literals continued
func mutatingMaps() {
	m := make(map[string]int)

	// Add the element
	m["Age"] = 32
	m["Experience"] = 11

	fmt.Println("The value for Age: ", m["Age"])

	// Change the element
	m["Age"] = 33
	fmt.Println("The value for Age: ", m["Age"])

	// Delete the element
	delete(m, "Age")

	// Test the element with a two-value assignment
	elem, ok := m["Age"]
	fmt.Println("The value for Age:", elem, "Present?", ok)

	elem, ok = m["Experience"]
	fmt.Println("The value for Experience:", elem, "Present?", ok)
}

func main() {
	maps()
	mapLiterals()
	mapLiteralsContinued()
	mutatingMaps()
}
