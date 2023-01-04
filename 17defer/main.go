package main

import "fmt"

func main() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Four")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v ", i)
	}
}
