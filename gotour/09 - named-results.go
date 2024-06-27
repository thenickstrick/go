// https://go.dev/tour/basics/7
package main

import "fmt"

// go's return values may be named. if so, they are treated as variables defined at the top of the function
// good names document the meaning of the return values
// a return statement without args returns the named return values; aka a "naked" return
// should be avoided in longer functions for readabilities' sake
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
