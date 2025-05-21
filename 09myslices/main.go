package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")
	var fruitList = []string{"Strawberry", "Pear"}
	fmt.Printf("The type is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Peach", "Apple")
	fmt.Println(fruitList)
	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)
	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)
	fruitList = append(fruitList[:3], fruitList...)
	fmt.Println(fruitList)
	highScores := make([]int, 4)
	highScores[0] = 783
	highScores[1] = 432
	highScores[2] = 987
	highScores[3] = 123
	// highScores[4]=777 --> this will throw an error out of bound as we have fixed the size to 4
	fmt.Println(highScores)
	highScores = append(highScores, 666, 444, 555)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	//remove value from slice based on index
	var courses = []string{"Web Dev", "App Dev", "Devops", "ML", "SDE", "Blockchain", "HFT"}
	var index int = 4
	// courses = append(courses[:index], courses[index+1:]...)
	courses = slices.Delete(courses, index, index+1)
	fmt.Println(courses)
}
