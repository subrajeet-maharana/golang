package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// fmt.Println("Welcome to the Maths in Golang")
	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5
	// fmt.Println("The sum is: ", myNumberOne+int(myNumberTwo))
	// //Here we lost the .5

	// //Generate random number
	// fmt.Println(rand.Intn(5)) //It will always generate one number only between 0 and 5
	// rand.Seed(4)
	// fmt.Println(rand.Intn(5))
	// rand.Seed(5048)
	// fmt.Println(rand.Intn(5))
	// //You get the idea that it depends on Seed value but seed is not random, So the solution is here -
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))
	// //To get from 1 to 6
	// fmt.Println(rand.Intn(5) + 1)

	//random number using crypto package in golang
	myNewRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myNewRandomNum)
}
