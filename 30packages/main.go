package main

import (
	"fmt"

	"github.com/subrajeet-maharana/myniceprogram/helpers"
)

func main() {
	fmt.Println("Packages")
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	myVar.TypeNumber = 34
	fmt.Println(myVar.TypeName, myVar.TypeNumber)
}
