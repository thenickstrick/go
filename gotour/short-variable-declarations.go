package main

import "fmt"

// inside a function the ":=" short assignment statement can be used in place of a var declaration with implicit type
// this is not available outside a function as every statement begins with a keyword
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
