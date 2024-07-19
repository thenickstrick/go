package main
import "log"
import "time"

var n int = 0

func main(){
start := time.Now()
for i := 0; i < 1000000; i++{
	n += 1
}
elapsed := time.Since(start)
log.Printf("Loop took %s", elapsed)
}
