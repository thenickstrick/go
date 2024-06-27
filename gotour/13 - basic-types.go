// https://go.dev/tour/basics/11
package main

import (
	"fmt"
	"math/cmplx"
)

// variable declarations can be "factored" into blocks, similar to import statements
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 121)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// these are Go's basic types:
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
	// the int, uint, and uintptr types are usually 32 bits wide on 32-bit systems
	// and 64 bits wide on 64-bit systems
	// when you need an integer value you should use int unless the'se a specific reason
	// to use a sized or unsigned integer type
// byte 
	// alias for uint8
// rune 
	// alias for int32
  // represents a Unicode code point
// float32 float64
// complex64 complex128
