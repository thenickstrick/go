package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) 
// in go, exported names must be capitalized to be "exported"
// so in the 'math.Pi' line, 'pi' would not be valid as it is an 'unexported' name
}
