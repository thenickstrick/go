package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var intNum int = 32767 
	// largest int16 bit integer. add 1 and you get an integer overflow to -32768
	// int will default to 32 or 64 depending on the system's architecture
	// you can use uint for unsigned ints that only store positive integers,
	// allowing those vars to store twice max positive values as their signed counterparts
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	// floats in go require specifying 32 or 64 bit
	// the more bits available, the higher the precision if you want it at the cost of memory
	// this can prevent expected values from being stored in memory
	// for example, the float value above would return 12345679.000000 if stored using float32
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)
	// you cannot perform arithmetic operations using variables of different types 
	// without first casting them to a common type

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	// integer division results in an integer that is rounded down
	// if you want to get the remainder, use % or modulo
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello World"
	fmt.Println(myString)
	// strings are initialized with double quotes or back quotes
	// double quotes can only be used on a single line for definition
	// with back quotes you can format the string directly (new lines, etm)
	// strings can be concatenated like "Hello" + " " + "World"

	fmt.Println(len("test"))
	// you can get the length of a string by using the built-in len function
	// but this calculates the number of bytes, not characters in the string
	// go uses utf-8 encoding by default, so characters outside the vanilla ASCII
	// character set are stored with more than a single byte
	
	fmt.Println(utf8.RuneCountInString("some non standard character"))
	// if you want to calculate the length of a string in the number of characters
	// import the built-in unicode/utf8 package and used RuneCountInString function
	// runes are a data type in go and represent characters 

	var myRune rune = 'a'
	fmt.Println(myRune)
	// using a single character in single quotes is evaluated as a rune

	var myBoolean bool = false
	fmt.Println(myBoolean)
	// bools are bools
	
	var intNum3 int
	fmt.Println(intNum3)
	// you don't have to declare values for variable, go has default values for most types
	// all types of ints, floats, and runes default to 0
	// string is empty
	// booleans is false

	myVar := "text"
	fmt.Println(myVar)
	// you can omit the type in variable declaration if you set the value during declaration
	// this is known as type inference, and can be taken a step further by using := and omitting var
	// just note it is good practice to add the type when it isn't obvious, especially when importing 
	// data from a function 

	var1, var2 := 1, 2
	fmt.Println(var1, var2)
	// you can initialize multiple variables at once in one line

	const myConst string = "const value"
	fmt.Println(myConst)
	// just about all of what applies to variables applies to constants
	// the value just can't be changed once it's created
	// and they have to be initialized with a value explicitly 
}