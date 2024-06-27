// https://go.dev/tour/basics/2
package main

// this could be done with import "package" line by line, but why type that out all the time? 
// use a "factored" import statement
import (
	"fmt"
	"math"
)

func main () {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
	// the "%g" is a "format specifier" and writes to standard output 
	// utilizing key, value structured data 
}
