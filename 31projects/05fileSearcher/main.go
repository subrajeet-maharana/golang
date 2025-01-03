package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SearchConfig struct {
  CaseSensitive bool
  FilePath string
  SearchTerm string
}

func searchFile(config SearchConfig) error {
  file, err := os.Open(config.FilePath)
  if err != nil {
    return fmt.Errorf("Failed to open the file: %v", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  lineNum := 0
  matchCount := 0

  for scanner.Scan() {
    lineNum++
    line := scanner.Text()

    var searchLine, searchTerm string
    
    if config.CaseSensitive {
      searchLine = line
      searchTerm = config.SearchTerm
    } else {
      searchLine = strings.ToLower(line)
      searchTerm = strings.ToLower(config.SearchTerm)
    }

    if strings.Contains(searchLine, searchTerm) {
      fmt.Printf("Line is: %d: %s\n", lineNum, line)
      matchCount++
    }
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("Error reading file: %v\n", err)
  }

  fmt.Printf("\nFound %d matches\n", matchCount)
  return nil
}

func main() {
  config := SearchConfig{}
  if len(os.Args) !=3 {
    fmt.Println("Usage <filepath> <search_term>")
    os.Exit(1)
  }

  config.FilePath = os.Args[1]
  config.SearchTerm = os.Args[2]
  config.CaseSensitive = true
  if os.Args[3]=="false" {
    config.CaseSensitive = false
  }

  searchFile(config)
  if err := searchFile(config); err != nil {
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    os.Exit(1)
  }
}
