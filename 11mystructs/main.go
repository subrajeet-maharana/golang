package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in Golang!")
	subrajeet := User{"Subrajeet", "subrajeet@go.dev", true, 22}
	fmt.Println(subrajeet)
	fmt.Printf("The details of Subrajeet are %+v\n", subrajeet)
	fmt.Printf("The name is %v and email is %v\n", subrajeet.Name, subrajeet.Email)
	fmt.Println(subrajeet.printName())
}

// No concept of inheritance; No Parent-Child relation also in Golang;

func (u *User) printName() string {
	return u.Name
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
