// every go program is made up of packages, starting with package main
package main

// packages are imported and can use paths. this will allow usage of the last element on the import path
// for example the std library math/rand is comprised of files that begin with the statement package rand
import (
	"fmt"
	"math/rand"
)

func main () {
	fmt.Println("My favorite number is", rand.Intn(10))
}
