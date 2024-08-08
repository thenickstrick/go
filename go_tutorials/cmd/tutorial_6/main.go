package main

import (
	"fmt"
)

type gasEngine struct{
	mpg uint16
	gallons uint16
//	owner
}
// structs can hold mixed types in the forms of fields defined by name
// the values of the fields are set to the default of the type

type electricEngine struct{
	mpkwh uint16
	kwh uint16
}

type owner struct{
	name string
}
// yo dawg, you can create a struct that is referenced by another struct
// so you can struct while you struct

// methods are just like functions, but have a tie-in to structs
// and have direct access to the struct instance itself
// the (e gasEngine) portion assigns this method to the gasEngin struct
// this gives it access to the fields as well as other methods assigned
// to the struct it is assigned to
func (e gasEngine) milesLeft() uint16 {
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint16 {
	return e.kwh*e.mpkwh
}

// this is what is considered an interface 
type engine interface{
	milesLeft() uint16
}

func canMakeIt(e engine, miles uint16){
	if miles<=e.milesLeft(){
		fmt.Println("You can make it there!")
	}else{
		fmt.Println("Need to fuel up first!")
	}
}

func main(){
	var myEngine gasEngine = gasEngine{25, 15}
// using the struct literal syntax to assign values 
// the field names can be omitted and the fields 
// will accept the values in order 

// the fields can also be assigned values directly 
//	myEngine.mpg = 20 
//	myEngine.gallons = 13

	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in the tank: %v\n", myEngine.milesLeft())
	canMakeIt(myEngine, 50)

// an anonymous struct, however this isn't not reusable
//	var myEngine2 = struct{
//		mpg uint16
//		gallons uint16
//	}{25,15}
//	fmt.Println(myEngine2.mpg, myEngine2.gallons)
}
