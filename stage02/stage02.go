package stage02

import (
	"fmt"
	"math"
	"math/cmplx"
)

// MaxAttempts
// Constants cannot be declared using the := syntax
// Constants can be character, string, boolean, or numeric values.
const MaxAttempts = 10

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
}
