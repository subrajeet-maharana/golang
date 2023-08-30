package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Sushanto",
		Breed: "German Shepherd",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Abhishek",
		Color:         "Black",
		NumberOfTeeth: 32,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("The animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "bark"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "ugh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
