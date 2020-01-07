package main

import "fmt"

// A struct is a collection of fields.

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// Basic stratucs
	vertex := Vertex{1, 2}

	vertex.X = 5
	vertex.Y = 10

	fmt.Println(Vertex{1, 2})
	fmt.Println(vertex)
	fmt.Println(vertex.X)

	// Pointers to structs
	// To access the field X of a struct when we have
	// the struct pointer p we could write (*p).X. However,
	// that notation is cumbersome, so the language permits us
	// instead to write just p.X, without the explicit dereference.

	pointer := &vertex

	// cumbersome notation
	(*pointer).X = 1e9

	// without the explicit dereference, the same
	pointer.Y = 1e3

	fmt.Printf("The vertex equels: %v\n", vertex)

	// Struct Literals
	fmt.Println(v1, p, v2, v3)
}
