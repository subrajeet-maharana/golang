package main

import "fmt"

func main() {
	greet()
	greetTwo()
	result := adder(5, 68)
	fmt.Println("The result is", result)
	proResult, myMessage := proAdder(5, 68, 78, 98, 387, 89)
	fmt.Println("The pro result is", proResult)
	fmt.Println("The pro message is", myMessage)
  add := func (a, b int) int {
    return a+b
  }
  anResult := add(546,-98)
  fmt.Println("The result from anonymous function is: ", anResult)
}
func greetTwo() {
	fmt.Println("Welcome Again!")
}
func greet() {
	fmt.Println("Hello and welcome!")
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// (valOne int, valTwo int) int --> this int is the data type of return keyword which is called function signatures
// How to receive any number of parameters use ...
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi! This is your proAdder function."
}
