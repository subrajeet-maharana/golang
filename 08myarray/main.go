package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Arrays")
	var fruitList [4]string
	fruitList[0] = "Banana"
	fruitList[2] = "Orange"
	fruitList[3] = "Lemon"
	fmt.Println("The Fruit List is ", fruitList)
	fmt.Println("The length of Fruit List is ", len(fruitList))
	var vegList = [5]string{"Potato", "Cauliflower", "Cabbage"}
	fmt.Println("The Veg List is ", vegList)
	fmt.Println("The length of Veg List is ", len(vegList))
}
