package main

import "fmt"

// var declarations can include initializers, one per variable
// when the initializers are present, type can be omitted as the variable will take the type of the initializer
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
