package main

import "fmt"
// var can declare a list of variables
// as in function arg lists, the type is same as last listed
var c, python, java bool

func main() {
	// var statements can be at package or functions levels
	var i int
	fmt.Println(i, c, python, java)
}
