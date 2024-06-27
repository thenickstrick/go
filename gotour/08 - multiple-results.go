// https://go.dev/tour/basics/6
package main

import "fmt"

// a function can return any number of results
// as shown by the swap function returning two strings
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
