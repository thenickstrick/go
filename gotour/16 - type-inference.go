package main

import "fmt"

func main() {
//when declaring a var without specifying an explicit type in go 
//either by using the := syntax or var = expression
//the variable's typ is inferred from the value on the right hand side
	v := 0.867 + 0.5i // change me!
	fmt.Printf("v is of type %T\n", v)
}

//var i int
//j := i 
//j is an int
//when the right hand side of the declaration is typed, the new var is of that same type
//however, when the right hand side contains an untyped numeric constant, the new var type
//might be an in, float64, or complex128 depending on the precision of the constant
//i := 42 is int
//f := 3.142 is float64
//g := 0.867 + 0.5i is complex128