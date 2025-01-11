package main

import "fmt"

func main() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Four")
	// myDefer()
  panicAndRecover()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v ", i)
	}
}

func panicAndRecover() {
  fmt.Println("Panic and Recover Demonstration: ")
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered from panic!", r)
    }
  }()

  fmt.Println("Before Panic: ")
  panic("Something went wrong...")
  fmt.Println("After Panic: ")
}
