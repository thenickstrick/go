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
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
// by omitting the array length value, a slice is created
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
// unlike arrays, with slices you can add values using the built-in append function
// this function takes the slice as the first element and the value you want to append 
// to the end of the slice as a second. It then returns the slice with the new element appended
// note: initially an array is allocated that can hold the exact amount of values given
// when the append function goes to add the new value, a check is done to see if the
// underlying array has enough room for the additional value. if it does not,
// a new array is created with enough capacity and all values are copied to the new array
// meaning the array is now stored in a completely different memory location

// you can print the length of an underlying array using the cap built-in function
// to output the capacity of the slice
// this change in cap is the array changed from [4,5,6] -> [4,5,6,7,*,*]
// the new length is just 4,5,6,7, and the cap is the whole []
// the * are considered out of range in the slices index

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
// you can append multiple values to a slice by using the spread operator

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)
// another way to create a slice via the make function
// you can specify the length of the slice [3] and the capacity of the slice [8]
// by default the capacity will just be the length of the slice
// it's generally considered good practice to determine the cap of the slice if
// you have a rough idea of how many values your slice will need to store 
// this helps from a performance perspective preventing slow downs due to array reallocation

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
// create a map to lookup a value by its key using the make function
// in this example the keys are strings and the values are unsigned int8

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
// you can initialize a map with values immediately
// and you can print the value of a map's item by specifying the key 
	fmt.Println(myMap2["Jason"])
// go will always return a value for a map if provided a key that does not exist
// go will return the default value of the type used for the map. 
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}
// maps in go can return an optional second value that is a boolean
// that is based on if the value is in the map or not
//	delete(myMap2, "Adam")
// go has a built-in delete function to provide a map and a key to be deleted

	for whomsta, sunLoops := range myMap2{
		fmt.Printf("Name: %v, Age: %v \n", whomsta, sunLoops)
	}
// in loops the range keyword can be used to iterate over arrays, slices, and maps
// note that when iterating over a map, no order is preserved, meaning that the order
// the values are output even in this small example can differ between program execution

	for index, value := range intArr{
		fmt.Printf("Index: %v, Value: %v \n", index, value)
	}

	for i:=0; i<10; i++{
		fmt.Println(i)
	}
// distinct parts of this for: initialization; condition; post
// there are other post options: i--, i+=10, i-=10, i*=10, i/=10
}
