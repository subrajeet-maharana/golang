package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the feedback on our Pizza: ")

	//comma ok | comma error syntax
	input1, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for your rating %s", input1)
	fmt.Printf("Rating is of type %T", input1)
}
