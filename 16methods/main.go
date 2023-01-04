package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in Golang!")
	subrajeet := User{"Subrajeet", "subrajeet@go.dev", true, 22}
	fmt.Println(subrajeet)
	fmt.Printf("The details of Subrajeet are %+v\n", subrajeet)
	fmt.Printf("The name is %v and email is %v\n", subrajeet.Name, subrajeet.Email)
	subrajeet.GetStatus()
	subrajeet.NewEmail()
	fmt.Printf("The name is %v and email is %v\n", subrajeet.Name, subrajeet.Email)
}

// No concept of inheritance; No Parent-Child relation also in Golang;

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("The new email is: ", u.Email)
}
