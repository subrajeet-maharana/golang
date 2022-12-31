package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Time Format of Golang")
	presentTime := time.Now()
	fmt.Println("The current time is: ", presentTime.Format("02-01-2006 15:04:05 Monday"))
	//This is how crazy it gets with golang , we have to remember the layout string in the format(""). How Sick?

	//we can create a time in golang
	createdDate := time.Date(2025, time.November, 30, 0, 0, 0, 0, time.Local)
	fmt.Println("The created date is: ", createdDate)
	fmt.Println("The created and formatted date is: ", createdDate.Format("02-01-2006 Monday"))

	//We can basically build executables for windows, linux , macos in any machine use
	//GOOS="windows" go build -- for windows
	//GOOS="darwin" go build -- for mac
	//go env for all details
}
