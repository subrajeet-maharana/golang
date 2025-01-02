package main

import (
	"fmt"
	"os"
)

func main() {
  if len(os.Args) !=3 {
    fmt.Println("Usage <filepath> <search_term>")
    os.Exit(1)
  }
  filepath := os.Args[1]

  file, err := os.Open(filepath)
  if err != nil {
    fmt.Println("Error opening file.")
    os.Exit(1)
  }

  defer file.Close()
}
