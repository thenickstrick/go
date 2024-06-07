// https://go.dev/tour/basics/13
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// some simple numeric conversions:
// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)

// but these can declared/converted in a more simplistic fashion:
// i := 42
// f := float(i)
// u := uint(f)
