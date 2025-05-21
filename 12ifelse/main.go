package main

import "fmt"

func main() {
	fmt.Println("Welcome to Control Statements")
	var result string
	loginCount := 23
	if loginCount > 10 {
		result = "Hey ROBOT!"
	} else if loginCount == 10 {
		result = "Exactly 10"
	} else {
		result = "Good"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num > 10 {
		fmt.Println("Good Score")
	} else {
		fmt.Println("Bad Score")
	}

	// Goto Statement
	i := 0
reprintGoto:
	if i <= 10 {
		fmt.Printf("%d ", i)
		i++
		goto reprintGoto
	}
}
