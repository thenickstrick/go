// goroutines are a way to launch multiple functions 
// and have them execute concurrently
// remember: concurrency != parallel execution

package main

import (
	"fmt"
	"time"
	"sync" // adding support for wait groups to allow concurrency and mutex usage
)

var m = sync.RWMutex{}
// var m = sync.Mutex{}
// add a mutex (mutual exclusion) as a helper for the goroutine 
// to make writing to the slice safer in a concurrent program
// standard mutex can be great, but will completely lock out
// other goroutines from accessing the results slice
// this is where a RW mutex can be beneficial as there is added
// Read lock/unlock
var wg = sync.WaitGroup{}
// adding a waitgroup to basically add a counter to spawned goroutines
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}
// create a slice to store the results returned from the dbCall


func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
	// whenever a new goroutine, add 1 to the counter 
		go dbCall(i)
	// add the go keyword in front of the function that you 
	// want to run concurrently... but there's more to it
	}
	wg.Wait()
// wait for the goroutine counter to go back down to zero
// signaling that all the tasks have completed
	fmt.Printf("Total execution time: %v\n", time.Since(t0))	
	fmt.Printf("The results are %v\n", results)
}

func dbCall(i int) {
	// simulating DB call delay
	var delay float32 = 2000
// if you remove the randomness from this and allowing the possibility of two or more 
// processes to write to the same memory locations simultaneously can lead to erratic 
// memory corruption, so you should avoid running code like this. 
// you should use a mutex - example above
	time.Sleep(time.Duration(delay)*time.Millisecond)
//	fmt.Println("The result from the database is:", dbData[i])
//	m.Lock()
// essentially what happens here is when a goroutine reached this lock method 
// a check is performed to see if lock has already been set by another goroutine
// if it has, the goroutine will wait here until the lock is released before 
// proceeding and locking access itself to perform its task(s)
// locking the access to the append here to establish a queue to write to results slice
// it is important to think through where lock methods are implemented
// as they can validly be put in multiple places, but it is likely to 
// tank concurrency as other steps in the process would be okay to run simultaneously
// but also stall a queue if goroutines were prevented from proceeding with execution
//	results = append(results, dbData[i])
// append the values returned from the fake dbCall to the results slice
//	m.Unlock()
// releasing the lock method once done so other goroutines can obtain a lock as needed
	save(dbData[i])
	log()
	wg.Done()
// call the done method at the end to decrement the counter for the goroutine
}


// adding support for the RWMutex to allow goroutines to read, only blocking when
// writes may potentially be happening, allowing more flexibility over standard lock
func save(result string){
	m.Lock()
	results = append (results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)	
	m.RUnlock()
}