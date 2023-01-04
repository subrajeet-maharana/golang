package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "This needs to go in to file."
	file, err := os.Create("./myGoFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myGoFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	// fmt.Println("The data inside the file is: \n", databyte)
	fmt.Println("The data inside the file is: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
