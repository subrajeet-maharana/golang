package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SearchConfig struct {
  FilePath string
  SearchTerm string
}

func searchFile(config SearchConfig) {
  file, err := os.Open(config.FilePath)
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
    if strings.Contains(line, config.SearchTerm) {
      fmt.Printf("Line is: %d: %s\n", lineNum, line)
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("Error reading file: %v\n", err)
  }
}

func main() {
  config := SearchConfig{}
  if len(os.Args) !=3 {
    fmt.Println("Usage <filepath> <search_term>")
    os.Exit(1)
  }
  config.FilePath = os.Args[1]
  config.SearchTerm = os.Args[2]
  searchFile(config)
}
