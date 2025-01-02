package main

import (
	"fmt"
	"os"
)

func main() {
  var filepath string
  fmt.Println("Enter the file path: ");
  fmt.Scanln(&filepath)
  file, err := os.Open(filepath)
  if err != nil {
    fmt.Println("Error opening file.")
    os.Exit(1)
  }

  defer file.Close()
}
