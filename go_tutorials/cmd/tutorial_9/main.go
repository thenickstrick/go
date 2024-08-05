// channels are a way to enable goroutines to pass data between each other
// they hold data in a thread safe way (avoiding data races when r/w memory)
// and listen for when data to be added or removed, and block code 
// execution until one of these events occur

package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int, 5) 
// without a number specified, the channel is unbuffered
// adding a buffer size allows the process function to close out 
// allowing the main function to continue to read through the channel
	go process(c)
// for loops support the use of channels in go, as seen 
// in this example with the use of the range function
	for i:= range c{
		fmt.Println(i)
		time.Sleep((time.Second*1))
	} 
// use the make function followed by the chan keyword to create a channel
// then provide the type for the data that the channel will hold
//	c <- 1
// use the <- operator to add data to a channel
// you can think of a channel as an underlying array in this case this 
// is an unbuffered channel that only has enough room for one value
//	var i = <- c 
// you can retrieve the value from a channel using similar syntax
// a good way to think about this is the first instance is dropping
// the value 1 into the channel, and this instance is popping the 
// value out of the channel, making the channel empty and the variable
// i now holds the value that was popped out of the channel
//	fmt.Println(i)
// note that this code will not run as is, as it results in a deadlock
// when writing to an unbuffered channel, the code will block until 
// another goroutine reads from the channel, so in effect it will block
// forever, unable to reach the line where the channel is read from
// luckily go's runtime is smart enough to detect this and will throw
// a deadlock error rather than the code hanging here forever
}

func process(c chan int){
	defer close(c)
// without closing out the channel in this loop, the process will deadlock
// eventually once the loop finishes at 5, so either close the channel at 
// the beginning like so, or use the close function at the end of the loop
// defer statements are nice in this scenario as they run the whatever is
// specfied right before the function exits
	for i:=0; i<5; i++{
	c <- i
	}
	fmt.Println("Exiting process")
//	close(c)
// to avoid a deadlock, add a function that takes in a channel and 
// writes a value to it to be called as a goroutine in main
}
