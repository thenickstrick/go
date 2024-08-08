// pointers are variables that store memory locations rather than values
// pointers still have a memory address themselves and store the value 
// at another address

package main
import "fmt"

func main(){
	var p *int32 = new(int32)
// p will hold the memory address ofn an int32
// with the pointer uninitialized, the pointer will not
// have an address assigned to it in which to store an int32
// however you can use the built-in new function to get back
// a free memory location that is 32 bits wide that p can 
// use to store and int32 value
// not using new to assign a memory address to a pointer
// creates what is known as a nil pointer exception
// so make sure that the pointer isn't nil before 
// trying to assign a value to it
// you can't get a value at a memory address 
// that doesn't exist this way
	var i int32
	fmt.Printf("The value p points to is: %v\n", *p)
// the *p is called dereferencing the pointer
// the default value of 0 for int 32 will return
// when you initialize a pointer with a memory location
// is zeros out that memory location, ie it sets the 
// value at that memory location to the default value of the type
	fmt.Printf("The value of i is: %v\n", i)
// to change the values stored at the memory location of a pointer
// the star notation must be used again to assign a value
	p = &i 
// using the & notation here causes p to reference the memory 
// address location of i, not the value 
	*p = 1
// so setting the value of p using the * notation here
// will also change the value of i
// but this seems backwards to me at this moment
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	var k int32 = 2
	i = k
// when using regular variables, when setting i = k here
// the program will copy the value of k into i's memory location

// an exception to this practice of copying non-pointer variables is when working with slices
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
// changing the value like this in slices causes the value to change in the same way as pointers
// this is because under the hood slices are just pointers to underlying arrays

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing1)
	var result [5]float64 = square(&thing1)
// remove the & prepending thing1
	fmt.Printf("The result is: %v\n", result)
	fmt.Printf("The value of thing1 is: %v\n", thing1)

}
// a few things with the thing1/thing2 square function:
// 1 - in the way it it written now, this is using pointers to 
// make the two consume less memory. However, this causes any changes 
// made in one array affect the values in the other as they reference
// the same memory addresses in this configuration
// 2 - if you follow the comments to change it to be two different 
// arrays, however, this causes the memory util for this program
// to effectively double in requirement since a memory 
// block has to be given to each array to allow the copy of values
func square(thing2 *[5]float64) [5]float64{
// remove the * prepending the 5
	fmt.Printf("The memory location of the thing2 array is: %p\n", thing2)
// prepend thing2 with &
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
// remove the * prepending thing2
}

