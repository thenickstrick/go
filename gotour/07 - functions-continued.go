// https://go.dev/tour/basics/5
package main

import "fmt"

// when two or more consecutive named function parameters share a type, you omit the type on all but the last
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
