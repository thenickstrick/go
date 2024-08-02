package main

import (
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main(){
	t0 := time.Now()
	for i:=0; i<1000; i++{
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
}

func count() {
	var res int
	for i:=0; i<100000000; i++{
		res+=1
	}
	wg.Done()
}