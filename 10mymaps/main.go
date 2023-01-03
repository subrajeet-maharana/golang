package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"
	fmt.Println("List of all languages", languages)
	fmt.Println("JS is short for", languages["JS"])
	delete(languages, "GO")
	fmt.Println("List of all languages", languages)
	//loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
