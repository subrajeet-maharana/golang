package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops!")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }
	// for j := range days {
	// 	fmt.Println(days[j])
	// }
	// for index, day := range days {
	// 	fmt.Printf("Index is %v and Value is %v\n", index, day)
	// }
	// for _, day := range days {
	// 	fmt.Printf("Day is %v\n", day)
	// }
	randomValue := 1
	for randomValue < 10 {
		if randomValue == 6 {
			goto brand
		}
		if randomValue == 2 {
			randomValue++
			continue
		} else if randomValue == 8 {
			break
		}
		fmt.Println("Value is", randomValue)
		randomValue++
	}
brand:
	fmt.Println("Welcome to The Analysers")
}
