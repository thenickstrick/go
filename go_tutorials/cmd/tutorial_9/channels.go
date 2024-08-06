package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main(){
	// create some channels to pass an alert to when price checks are within specified range
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	// create slice of websites that are selling tofu/chicken so we can check their prices
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}	
	// spawn a goroutine for each website from the slice above that calls a price check 
	// function for both chicken and tofu on each website we want to check
	// and pass the website and channel to each function
	for i:= range websites{
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	// pass the value of the channels to the send message function 
	sendMessage(chickenChannel, tofuChannel)
}

// both of these price check functions are doing the same things:
// accepting the website being the price check is being done against
// and checking the price every second, then once the price threshold 
// is reached, set the value of the channel to the name of the website
// where the price crossed the threshold, then breaks out of the goroutine
func checkTofuPrices(website string, tofuChannel chan string){
	for {
		time.Sleep(time.Second*1)
		var tofuPrice = rand.Float32()*20
		if tofuPrice<MAX_TOFU_PRICE{
			tofuChannel <- website
			break
		}
	}
}

func checkChickenPrices(website string, chickenChannel chan string){
	for {
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice<=MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}

// accepts the values of the channels from the main function, and 
// based on if it's the chicken or tofu channel, outputs message
// for the specified channel and website where the price is below
// the desired threshold
func sendMessage(chickenChannel chan string, tofuChannel chan string){
	select{
		case website := <-chickenChannel:
			fmt.Printf("Text Sent: Found a deal on chicken at %v\n", website)
		case website := <-tofuChannel:
			fmt.Printf("Email Sent: Found a deal on tofu at %v\n", website)
		}
}
