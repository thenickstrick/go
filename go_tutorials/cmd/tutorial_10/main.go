package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type gasEngine struct{
	gallons float32
	mpg float32
}

type electricEngine struct{
	kwh float32
	mpkwh float32
}

// using generics in structs like this allows you to support
// instantiating differing types of cars that use different 
// engine types defined above
type car [T gasEngine | electricEngine] struct{
	carMake string
	carModel string
	engine T
}

type contactInfo struct{
	Name string
	Email string
}

type purchaseInfo struct{
	Name string
	Price float32
	Amount int
}

func main(){
	// in these cases the [type] could be omitted and the 
	// type of values in the slice can be inferred, but this
	// is not always a good fit 
	var intSlice = []int{}
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(isEmpty[int](intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))
	fmt.Println(isEmpty[float32](float32Slice))

	var gasCar = car[gasEngine]{
		carMake: "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg: 40,
		},
	}
	var electricCar = car[electricEngine]{
		carMake: "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh: 57.5, 
			mpkwh: 4.17,
		},
	}

	// in these struct examples and the associated function,
	// if the struct type is not passed in the function won't
	// know what struct to unmarshall the json string into
	// var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	// fmt.Printf("%+v\n", contacts)

	// var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	// fmt.Printf("%+v\n", purchases)
}

// setting up a generic to accept additional parameters
// that are defined within the brackets of the func
// except instead of passing in a value, a type is used
// essentially the type is subbed in for T
// the "any" type is an option for generics, however in 
// this case it is not usable since any is not compatible 
// with the addition operator
func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum+= v
	}
	return sum
}

func isEmpty[T any](slice []T) bool{
	return len(slice)==0
}

// func loadJSON[T contactInfo | purchaseInfo](filepath string) []T{
// 	data, _ = ioutil.ReadFile(filepath)

// 	var loaded = []T{}
// 	json.Unmarshal(data, &loaded)
// 	// unmarshal is essentially loading the json into a struct 
// 	return loaded
// }
