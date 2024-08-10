package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = "résumé"
	var indexed = myString[1]
// you can index the string just like arrays using the same notation
	fmt.Println(indexed) 
// printing out the indexed value returns 114, hm
	fmt.Printf("%v, %T\n", indexed, indexed) 
// printing out the type using %T
// iterate over the index to see why the value is 114 and type is uint8
	for i, v := range myString{ 
		fmt.Println(i, v)
	}
// outputs the following:
// 0 114
// 1 233
// 3 115
// 4 117
// 5 109
// 6 233
// showing that go defaults to UTF-8 encoding
// the 114 is the byte value of the letter r that resides at index 0
// printing the index for 1 prints the first half of the UTF-8 encoding 
// of the é character. This also shows that the range keyword is performing
// some additional work, providing both bytes for the é character instead 
// of the first half. this is why index 2 is not shown when iterating as 
// it is used for the 2nd byte for the é character

	fmt.Printf("The length of 'myString' is %v\n", len(myString))
// shows the number of bytes of the string and not the number of characters
	var myString2 = []rune("résumé") 
// casting the string to an array of runes allows strings to be interacted
// with in a more straight-forward fashion. not dealing with the underlying 
// byte array etm
// runs are essentially unicode point numbers that represent the character
// and are an alias for int32 
	var indexed2 = myString2[1]
	fmt.Println(indexed2) 
	fmt.Printf("%v, %T\n", indexed2, indexed2) 
	for i, v := range myString2{ 
		fmt.Println(i, v)
	}
	fmt.Printf("The length of 'myString' is %v\n", len(myString2))

	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)
// a rune type can be declared using a single quote

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""
	for i := range strSlice{
		catStr += strSlice[i]
	}
	fmt.Printf("%v\n", catStr)
// strings can be concatenated together, and I already subscribed Alex!
// strings are immutable in go, meaning the cannot be modified once created
// in this example the concatenation to create this string and store it in 
// a variable is pretty inefficient, since it creates a new string every time
// instead you can import go's built-in strings package and use a string builder
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	// instead of the + operator, using WriteString method to append values 
	// and only at the end is the new string created from the appended values
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("%v\n", catStr2)
}
