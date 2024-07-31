package main

import "fmt"

func main(){
	var intArr [3]int32 = [3]int32{1,2,3}
// the square bracket indicates a fixed number of elements defined for this array
// the length of an array declared this way cannot be changed after it is initialized
// and the type defined specifies the type for every element in the array
// the default values for each element will be the same typical type defaults
// since the array is 4 elements and int32 consumes 4 bytes of memory
// the array will use a contiguous 12 bytes of memory when initialized
// arrays can also be initialized with variables as well as inferred with :=
// and can be inferred even further by [...]int32{1,2,3} and still be a fixed size of 3
	intArr[1] = 123
// change the value of index of 1
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
// indexable, with a typical zero index structure
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
// the memory location of each element can be printed by including an & with the var

	
// according to the go doc here: https://go.dev/doc/effective_go#slices
// "slices wrap arrays to give a more general, powerful, and convenient interface to 
// sequences of data. Except for items with explicit dimension such as transformation 
// matrices, most array programming in Go is done with slices rather than simple arrays."
// essentially slices in go are just arrays with some additional functionality.
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
// by omitting the array length value, a slice is created
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
// unlike arrays, with slices you can add values using the built-in append function
// this function takes the slice as the first element and the value you want to append 
// to the end of the slice as a second. It then returns the slice with the new element appended
// note: initially an array is allocated that can hold the exact amount of values given
// when the append function goes to add the new value, a check is done to see if the
// underlying array has enough room for the additional value. if it does not,
// a new array is created with enough capacity and all values are copied to the new array
// meaning the array is now stored in a completely different memory location
}