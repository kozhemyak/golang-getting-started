package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// MaxAttempts
// Constants cannot be declared using the := syntax
// Constants can be character, string, boolean, or numeric values.
const MaxAttempts = 10

// "factored" into block
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

// Variables declared without an explicit initial value are given their zero value.

// Types:
// bool
// string
// int (32, 64)  - int8  int16  int32  int64
// uint (32, 64) - uint8 uint16 uint32 uint64 uintptr
// byte	- alias for uint8
// rune - alias for int32
// float32 float64
// complex64 complex128

// "factored" into block
var (
	ToBe    bool       = false
	Maint   uint64     = 1<<64 - 1
	Complex complex128 = cmplx.Sqrt(-5 + 12i)
)

// Run all this calculations
func Start() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", Maint, Maint)
	fmt.Printf("Type: %T Value: %v\n", Complex, Complex)

	// Type conversions: T(v)
	f := math.Sqrt(float64(3) / float64(4))
	z := uint(f)

	fmt.Printf("Type: %T Value: %f\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Type inference
	v := 4.0/3 + MaxAttempts // change me!
	fmt.Printf("v is of type %T\n", v)

	// Contsts
	// Numeric constants are high-precision values.
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func main() {
	Start()
}
