package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers!")
	var ptr *int
	fmt.Println("The value of pointer is ", ptr)
	myNumber := 34
	var ptr1 = &myNumber
	fmt.Println("The address of pointer is", ptr1)
	fmt.Println("The value of pointer is", *ptr1)
	*ptr1 = *ptr1 * 5
	fmt.Println("The new value is ", myNumber)
}
