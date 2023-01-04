package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch Case in Golang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The value of the dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You have got 1 and can open now.")
	case 2:
		fmt.Println("You can move 2 spots.")
	case 3:
		fmt.Println("You can move 3 spots.")
	case 4:
		fmt.Println("You can move 4 spots.")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spots.")
	case 6:
		fmt.Println("You can move 6 spots and roll the dice again.")
	default:
		fmt.Println("What was that!")
	}
}
