// https://go.dev/tour/basics/4
package main

import "fmt"

// a function can take zero or more arguments
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
} 
