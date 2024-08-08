package main

import ( 
	"errors"
	"fmt"
)

func main(){
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	if err!=nil{
		fmt.Printf(err.Error())
	}else if remainder == 0{
		fmt.Printf("The result of the integer division is %v, result\n")
	}else{
	fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}
// the above if else can be rewritten as a switch statement
  switch{
		case err!=nil:
			fmt.Printf(err.Error())
		case remainder==0:
			fmt.Printf("The result of the integer division is %v", result)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}
	switch remainder{
	case 0:
		fmt.Printf("The division was exact")
	case 1,2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
}

func printMe(printValue string){
// the parentheses at the start of functions allow the passing in of parameters 
// in this example an input value with a parameter name of printValue of type string
// is defined to accept the variable defined in the main function to be passed in
// the parameter type is enforced, so an int can't be passed in in place of a string
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
// when you have a function that returns some value back to the caller
// you have to specify the type that it is returning, and this can 
// also return multiple values by parenthesizing the returned types 
// and supplying multiple types: ) (int, int) { 
// instead of ) int {
	var err error
// error is a built-in type in go, who's default value is nil
	if denominator==0{
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
// return the results to make them accessible to the caller
}
