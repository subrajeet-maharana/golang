package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
  if len(os.Args) !=3 {
    fmt.Println("Usage <filepath> <search_term>")
    os.Exit(1)
  }
  filepath := os.Args[1]
  searchTerm := os.Args[2]

  file, err := os.Open(filepath)
  if err != nil {
    fmt.Println("Error opening file.")
    os.Exit(1)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  lineNum := 0

  for scanner.Scan() {
    lineNum++
    line := scanner.Text()
    if strings.Contains(line, searchTerm) {
      fmt.Printf("Line is: %d: %s\n", lineNum, line)
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("Error reading file: %v\n", err)
  }
}
